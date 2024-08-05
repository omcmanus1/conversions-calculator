"use client";

import Link from "next/link";
import { Button } from "@/components/ui/button";

export default function Page() {
  return (
    <Button variant="outline" asChild>
      <Link href="/recipes/usa">Don't Click Me</Link>
    </Button>
  );
}
