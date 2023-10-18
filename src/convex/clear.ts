import { mutation } from "./_generated/server";
import { v } from "convex/values";

export default mutation({
  args: { token: v.string() },
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

