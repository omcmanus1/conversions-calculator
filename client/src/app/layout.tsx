"use client";

import { Button } from "@/components/ui/button";
import "./globals.css";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Link from "next/link";
import Navigation from "@/components/navigation";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const navigationTabs = [
    { title: "Home", path: "/home" },
    { title: "Convert US To Metric", path: "/convert-usa" },
    { title: "Convert Metric To US", path: "/convert-metric" },
  ];
  return (
    <html lang="en">
      <head>
        <title>Recipe Unit Converter</title>
      </head>
      <body className={`flex flex-col items-center bg-slate-200 ${inter.className}`}>
        <header className="m-5 flex flex-col items-center">
          <h1 className="m-5 text-6xl uppercase">Recipe Unit Converter</h1>
          <Navigation navigationTabs={navigationTabs}></Navigation>
          {/* <nav>
            <Button variant="outline" className="mr-2" asChild>
              <Link href="/">Home</Link>
            </Button>
            <Button variant="outline" className="mr-2" asChild>
              <Link href="/convert-usa">Convert US To Metric</Link>
            </Button>
            <Button variant="outline" className="mr-2" asChild>
              <Link href="/convert-metric">Convert Metric To US</Link>
            </Button>
          </nav> */}
        </header>
        <div>{children}</div>
      </body>
    </html>
  );
}
