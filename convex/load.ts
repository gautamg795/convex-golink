import { query } from "./_generated/server";
import { Document } from "./_generated/dataModel";

export const loadOne = query(
  async ({ db }, normalizedId: string): Promise<Document<"links"> | null> => {
    return await db
      .query("links")
      .withIndex("by_normalizedId", (q) => q.eq("normalizedId", normalizedId))
      .first();
  }
);

export const loadAll = query(
  async ({ db }): Promise<Array<Document<"links">>> => {
    // TODO: Paginate
    return await db.query("links").fullTableScan().collect();
  }
);
