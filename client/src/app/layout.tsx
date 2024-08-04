"use client";

import "./globals.css";
import { Inter } from "next/font/google";
import { usePathname } from "next/navigation";
import Navigation from "@/components/navigation";
import { Analytics } from "@vercel/analytics/react";
import { Toaster } from "@/components/ui/sonner";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({ children }: { children: React.ReactNode }) {
  const topLevelTabs = [
    { title: "Recipes", path: "/recipes" },
    { title: "Measurements", path: "/measurements" },
  ];

  const pathname = usePathname();
  const rootSite = pathname.includes("recipe") ? "Recipe Unit" : "Measurement";

  return (
    <html lang="en">
      <head className="text-center">
        <title>Unit Converter</title>
      </head>
      <body
        className={`flex flex-col min-h-screen items-center bg-slate-200 text-center ${inter.className}`}
      >
        <header className="md:m-5 flex flex-col items-center">
          <Navigation tabs={topLevelTabs}></Navigation>
          <h1 className="m-4 text-4xl md:text-6xl uppercase shrink">
            {rootSite} Converter
          </h1>
        </header>
        <div>{children}</div>
        <Toaster richColors />
        <footer className="text-center p-1 absolute bottom-0 bg-slate-200">
          <a
            href="https://www.flaticon.com/free-icons/change"
            title="change icons"
            className="text-slate-400 text-xs"
          >
            Change icons created by Smashicons - Flaticon
          </a>
          <Analytics />
        </footer>
      </body>
    </html>
  );
}
