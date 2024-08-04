"use client";

import Navigation from "@/components/navigation";
import { usePathname } from "next/navigation";

export default function TopLevelNav() {
  const topLevelTabs = [
    { title: "Recipes", path: "/recipes" },
    { title: "Measurements", path: "/measurements" },
  ];

  const pathname = usePathname();
  const chosenSite = pathname.includes("recipe") ? "Recipe Unit" : "Measurement";

  return (
    <>
      <Navigation tabs={topLevelTabs}></Navigation>
      <h1 className="m-4 text-4xl md:text-6xl uppercase shrink">
        {chosenSite} Converter
      </h1>
    </>
  );
}
