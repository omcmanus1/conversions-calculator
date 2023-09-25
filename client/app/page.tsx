"use client";

import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function Page() {
  return (
    <Button variant="outline" asChild>
      <Link href="/convert-usa">Don't Click Me</Link>
    </Button>
  );
}