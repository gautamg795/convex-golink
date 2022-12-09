import { mutation } from "./_generated/server";
import { Document } from "./_generated/dataModel";

export default mutation(async ({ db }, link: Document<"links">) => {
  let existing = await db
    .query("links")
    .withIndex("by_normalizedId", (q) =>
      q.eq("normalizedId", link.normalizedId)
    )
    .first();
  if (existing !== null) {
    await db.replace(existing._id, link);
    return;
  }
  await db.insert("links", link);
});
