import "./globals.css";
import { Inter } from "next/font/google";
import { Analytics } from "@vercel/analytics/react";
import { Toaster } from "@/components/ui/sonner";
import TopLevelNav from "./top-level-nav";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <head className="text-center">
        <title>Conversions Calculator</title>
      </head>
      <body
        className={`flex flex-col min-h-screen items-center bg-slate-200 text-center ${inter.className}`}
      >
        <header className="md:m-5 flex flex-col items-center">
          <TopLevelNav />
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
