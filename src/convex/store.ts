import { mutation } from "./_generated/server";
import { v } from "convex/values";
import { LinkDoc } from "./schema";

export default mutation({
  args: { link: v.object(LinkDoc), token: v.string() },
  handler: async (ctx, { link, token }) => {
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
      throw new Error("Invalid authorization token");
    }
    let existing = await ctx.db
      .query("links")
      .withIndex("by_normalizedId", (q) =>
        q.eq("normalizedId", link.normalizedId)
      )
      .first();
    if (existing !== null) {
      await ctx.db.replace(existing._id, link);
      return;
    }
    await ctx.db.insert("links", link);
  },
});
