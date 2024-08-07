"use client";

import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu";
import { cn } from "@/utils/shadutils";
import Link from "next/link";
import { usePathname } from "next/navigation";
import TooltipSh from "../tooltip";

type NavigationTab = {
  path: string;
  title: string;
  disabled?: boolean;
  tooltip?: string;
};

export default function Navigation({ tabs }: { tabs: NavigationTab[] }) {
  const pathname = usePathname() || "";

  return (
    <NavigationMenu className="mb-2">
      <NavigationMenuList className="flex-col md:flex-row">
        {tabs.map((tab) => {
          const isSelected = pathname.includes(tab.path);
          return (
            <NavigationMenuItem key={`${tab.path}_key`} className="mt-2">
              {tab.disabled ? (
                <div
                  className={cn(
                    navigationMenuTriggerStyle(),
                    isSelected && "!bg-teal-100",
                    "border border-blue-300",
                    tab.disabled && "opacity-50 cursor-not-allowed"
                  )}
                >
                  <TooltipSh title={tab.title} tooltip={tab?.tooltip || ""} />
                </div>
              ) : (
                <Link href={tab.path} legacyBehavior passHref>
                  <NavigationMenuLink
                    className={cn(
                      navigationMenuTriggerStyle(),
                      isSelected && "!bg-teal-100",
                      "border border-blue-300",
                      tab.disabled && "cursor-not-allowed opacity-50"
                    )}
                  >
                    {tab.title}
                  </NavigationMenuLink>
                </Link>
              )}
            </NavigationMenuItem>
          );
        })}
      </NavigationMenuList>
    </NavigationMenu>
  );
}
