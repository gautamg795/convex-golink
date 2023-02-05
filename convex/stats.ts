import { query, mutation } from "./_generated/server";
import { Document, Id } from "./_generated/dataModel";

export const loadStats = query(async ({ db }, token: string) => {
  if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
    throw new Error("Invalid authorization token");
  }
  let stats = new Map<string, number>();
  for await (const link of db.query("links").fullTableScan()) {
    const clicks = (
      await db
        .query("stats")
        .withIndex("byLink", (q) => q.eq("link", link._id))
        .first()
    )?.clicks;
    if (clicks) {
      stats.set(link.normalizedId, clicks);
    }
  }
  console.log(stats);
  return stats;
});

export const saveStats = mutation(async ({ db }, stats: Object, token: string) => {
  if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
    throw new Error("Invalid authorization token");
  }
  for (const [normalizedId, clicks] of Object.entries(stats)) {
    const link = await db
      .query("links")
      .withIndex("by_normalizedId", (q) => q.eq("normalizedId", normalizedId))
      .first();
    if (link !== null) {
      let stat = await db
        .query("stats")
        .withIndex("byLink", (q) => q.eq("link", link._id))
        .first();
      if (stat !== null) {
        stat.clicks += clicks;
        await db.replace(stat._id, stat);
      } else {
        await db.insert("stats", { link: link._id, clicks: clicks });
      }
    } else {
      console.warn("Writing stats for nonexistent link: ", normalizedId);
    }
  }
});
