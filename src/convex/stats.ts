import { query, mutation } from "./_generated/server";
import { v } from "convex/values";

export const loadStats = query({
  args: { token: v.string() },
  returns: v.any(),
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
        stats[link.normalizedId] = clicks;
      }
    }
    return stats;
  },
});

export const saveStats = mutation({
  args: { stats: v.record(v.string(), v.number()), token: v.string() },
  returns: v.null(),
  handler: async (ctx, { stats, token }) => {
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
      throw new Error("Invalid authorization token");
    }
    const typedStats: Record<string, number> = stats;
    for (const [normalizedId, clicks] of Object.entries(typedStats)) {
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
