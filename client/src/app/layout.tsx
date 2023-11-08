"use client";

import "./globals.css";
import { Inter } from "next/font/google";
import Navigation from "@/components/navigation";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({ children }: { children: React.ReactNode }) {
  const navigationTabs = [
    { title: "Convert US To Metric", path: "/recipes/convert-usa" },
    { title: "Convert Metric To US", path: "/recipes/convert-metric" },
    { title: "Convert List", path: "/recipes/convert-list" },
  ];
  return (
    <html lang="en">
      <head className="text-center">
        <title>Recipe Unit Converter</title>
      </head>
      <body
        className={`flex flex-col min-h-screen items-center bg-slate-200 text-center ${inter.className}`}
      >
        <header className="m-5 flex flex-col items-center">
          <h1 className="m-4 text-4xl md:text-6xl uppercase shrink">
            Recipe Unit Converter
          </h1>
          <Navigation navigationTabs={navigationTabs}></Navigation>
        </header>
        <div>{children}</div>
        <footer className="text-center p-1 absolute bottom-0 bg-slate-200">
          <a
            href="https://www.flaticon.com/free-icons/change"
            title="change icons"
            className="text-slate-400 text-xs"
          >
            Change icons created by Smashicons - Flaticon
          </a>
        </footer>
      </body>
    </html>
  );
}
