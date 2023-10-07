"use client";

import "./globals.css";
import { Inter } from "next/font/google";
import Navigation from "@/components/navigation";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const navigationTabs = [
    // { title: "Home", path: "/home" },
    { title: "Convert US To Metric", path: "/recipes/convert-usa" },
    { title: "Convert Metric To US", path: "/recipes/convert-metric" },
  ];
  return (
    <html lang="en">
      <head>
        <title>Recipe Unit Converter</title>
      </head>
      <body
        className={`flex flex-col items-center bg-slate-200 ${inter.className}`}
      >
        <header className="m-5 flex flex-col items-center">
          <h1 className="m-5 text-6xl uppercase">Recipe Unit Converter</h1>
          <Navigation navigationTabs={navigationTabs}></Navigation>
        </header>
        <div>{children}</div>
      </body>
    </html>
  );
}
