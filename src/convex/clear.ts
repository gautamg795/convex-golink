import { mutation } from "./_generated/server";
import { v } from "convex/values";

export default mutation({
  args: { token: v.string() },
  returns: v.null(),
  handler: async (ctx, { token }) => {
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
      throw new Error("Invalid authorization token");
    }
    let deletions = [];
    for await (const link of ctx.db.query("links").fullTableScan()) {
      deletions.push(ctx.db.delete(link._id));
    }
    for await (const stat of ctx.db.query("stats").fullTableScan()) {
      deletions.push(ctx.db.delete(stat._id));
    }
    await Promise.all(deletions);
  },
});

export const deleteOne = mutation({
  args: { normalizedId: v.string(), token: v.string() },
  returns: v.null(),
  handler: async (ctx, { normalizedId, token }) => {
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
      throw new Error("Invalid authorization token");
    }
    const link = await ctx.db
      .query("links")
      .withIndex("by_normalizedId", (q) => q.eq("normalizedId", normalizedId))
      .first();

    if (link !== null) {
      let stats = await ctx.db.query("stats").withIndex("byLink", (q) => q.eq("link", link._id)).collect();
      for (const stat of stats) {
        await ctx.db.delete(stat._id);
      }
      await ctx.db.delete(link._id);
    }
  },
});

