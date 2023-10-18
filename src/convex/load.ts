import { query } from "./_generated/server";
import { v } from "convex/values";

export const loadOne = query({
  args: { normalizedId: v.string(), token: v.string() },
  handler: async (ctx, { normalizedId, token }) => {
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
      throw new Error("Invalid authorization token");
    }
    return await ctx.db
      .query("links")
      .withIndex("by_normalizedId", (q) => q.eq("normalizedId", normalizedId))
      .first();
  },
});

export const loadAll = query({
  args: { token: v.string() },
  handler: async (ctx, { token }) => {
    // TODO: Paginate
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
      throw new Error("Invalid authorization token");
    }
    return await ctx.db.query("links").fullTableScan().collect();
  },
});
