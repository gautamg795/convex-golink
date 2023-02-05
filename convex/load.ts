import { query } from "./_generated/server";
import { Document } from "./_generated/dataModel";

export const loadOne = query(
  async ({ db }, normalizedId: string, token: string): Promise<Document<"links"> | null> => {
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
      throw new Error("Invalid authorization token");
    }
    return await db
      .query("links")
      .withIndex("by_normalizedId", (q) => q.eq("normalizedId", normalizedId))
      .first();
  }
);

export const loadAll = query(
  async ({ db }, token: string): Promise<Array<Document<"links">>> => {
    // TODO: Paginate
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
      throw new Error("Invalid authorization token");
    }
    return await db.query("links").fullTableScan().collect();
  }
);
