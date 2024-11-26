import { v } from "convex/values";
import { mutation } from "./_generated/server";

export default mutation({
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
      throw new Error("Can't delete nonexistent link with id " + normalizedId);
    }
    await ctx.db.delete(existing._id);
  },
});
