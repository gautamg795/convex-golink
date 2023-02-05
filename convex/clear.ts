import { mutation } from "./_generated/server";

export default mutation(async ({ db }, token: string) => {
    if (token === "" || token !== process.env.CONVEX_AUTH_TOKEN) {
        throw new Error("Invalid authorization token");
    }
    let deletions = [];
    for await (const link of db.query("links").fullTableScan()) {
        deletions.push(db.delete(link._id));
    }
    for await (const stat of db.query("stats").fullTableScan()) {
        deletions.push(db.delete(stat._id));
    }
    await Promise.all(deletions);
});