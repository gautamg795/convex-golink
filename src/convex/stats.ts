import { query, mutation } from "./_generated/server";
import { v } from "convex/values";

export const loadStats = query({
  args: { token: v.string() },
  handler: async (ctx, { token }) => {
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
      throw new Error("Invalid authorization token");
    }
    let stats: Record<string, number> = {};
    for await (const link of ctx.db.query("links").fullTableScan()) {
      const clicks = (
        await ctx.db
          .query("stats")
          .withIndex("byLink", (q) => q.eq("link", link._id))
          .first()
      )?.clicks;
      if (clicks) {
        stats[link.short] = clicks;
      }
    }
    return stats;
  },
});

export const saveStats = mutation({
  args: { stats: v.record(v.string(), v.number()), token: v.string() },
  handler: async (ctx, { stats, token }) => {
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
      throw new Error("Invalid authorization token");
    }
    for (const [normalizedId, clicks] of Object.entries(stats)) {
      const link = await ctx.db
        .query("links")
        .withIndex("by_normalizedId", (q) => q.eq("normalizedId", normalizedId))
        .first();
      if (link !== null) {
        let stat = await ctx.db
          .query("stats")
          .withIndex("byLink", (q) => q.eq("link", link._id))
          .first();
        if (stat !== null) {
          stat.clicks += clicks;
          await ctx.db.replace(stat._id, stat);
        } else {
          await ctx.db.insert("stats", { link: link._id, clicks: clicks });
        }
      } else {
        console.warn("Writing stats for nonexistent link: ", normalizedId);
      }
    }
  },
});

export const deleteStats = mutation({
  args: { normalizedId: v.string(), token: v.string() },
  handler: async (ctx, { normalizedId, token }) => {
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
      throw new Error("Invalid authorization token");
    }
    let existing = await ctx.db
      .query("links")
      .withIndex("by_normalizedId", (q) => q.eq("normalizedId", normalizedId))
      .first();
    if (existing === null) {
      return;
    }
    let stats = await ctx.db
      .query("stats")
      .withIndex("byLink", (q) => q.eq("link", existing!._id))
      .first();
    if (stats !== null) {
      await ctx.db.delete(stats._id);
    }
  },
});
