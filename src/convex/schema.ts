import { defineSchema, defineTable, s } from "convex/schema";

export default defineSchema({
  links: defineTable({
    normalizedId: s.string(),
    short: s.string(),
    long: s.string(),
    created: s.number(),
    lastEdit: s.number(),
    owner: s.string(),
  }).index("by_normalizedId", ["normalizedId"]),
  stats: defineTable({
    link: s.id("links"),
    clicks: s.number(),
  }).index("byLink", ["link"]),
});
