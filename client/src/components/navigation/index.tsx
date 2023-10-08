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

export default function Navigation({
  navigationTabs,
}: {
  navigationTabs: Array<NavigationTab>;
}) {
  const pathname = usePathname() || "";

  return (
    <NavigationMenu className="mb-5">
      <NavigationMenuList>
        {navigationTabs.map((tab) => {
          const isSelected = pathname.includes(tab.path);
          return (
            <NavigationMenuItem key={`${tab.path}_key`}>
              <Button className="mr-2" asChild>
                <Link href={tab.path} legacyBehavior passHref>
                  <NavigationMenuLink
                    className={cn(
                      navigationMenuTriggerStyle(),
                      isSelected && "underline",
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
