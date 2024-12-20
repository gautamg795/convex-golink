import { defineSchema, defineTable } from "convex/server";
import { v } from "convex/values";

export const LinkDoc = {
  normalizedId: v.string(),
  short: v.string(),
  long: v.string(),
  created: v.number(),
  lastEdit: v.number(),
  owner: v.string(),
};

export default defineSchema({
  links: defineTable(LinkDoc).index("by_normalizedId", ["normalizedId"]),
  stats: defineTable({
    link: v.id("links"),
    clicks: v.number(),
  }).index("byLink", ["link"]),
});
