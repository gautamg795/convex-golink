import { query } from "./_generated/server";
import { v } from "convex/values";
import { LinkTable } from "./schema";

// Combine validators with system fields
export const loadOne = query({
  args: { normalizedId: v.string(), token: v.string() },
  returns: v.union(v.object({ ...LinkTable.validator.fields, _creationTime: v.number(), _id: v.id("links") }), v.null()),
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
  returns: v.array(v.object({ ...LinkTable.validator.fields, _creationTime: v.number(), _id: v.id("links") })),
  handler: async (ctx, { token }) => {
    // TODO: Paginate
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
      throw new Error("Invalid authorization token");
    }
    return await ctx.db.query("links").fullTableScan().collect();
  },
});
