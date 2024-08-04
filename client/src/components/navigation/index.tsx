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
import { Button } from "@/components/ui/button";

type NavigationTab = {
  path: string;
  title: string;
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
              <Button className="mr-2" asChild>
                <Link href={tab.path} legacyBehavior passHref>
                  <NavigationMenuLink
                    className={cn(
                      navigationMenuTriggerStyle(),
                      isSelected && "underline bg-teal-100",
                      "border border-blue-300"
                    )}
                  >
                    {tab.title}
                  </NavigationMenuLink>
                </Link>
              </Button>
            </NavigationMenuItem>
          );
        })}
      </NavigationMenuList>
    </NavigationMenu>
  );
}
