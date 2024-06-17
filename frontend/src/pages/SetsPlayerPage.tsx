import * as React from "react";

import { Card, CardContent } from "@/components/ui/card";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";

export function SetsPlayerPage() {
  return (
    <Carousel
      opts={{
        align: "start",
      }}
      orientation="vertical"
      className="h-screen"
    >
      <CarouselContent className="-mt-1">
        {Array.from({ length: 20 }).map((_, index) => (
          <CarouselItem key={index} className="">
            <div className="p-1">
              <Card>
                <CardContent className="flex items-center justify-center p-6">
                  <span className="text-3xl font-semibold">{index + 1}</span>
                </CardContent>
              </Card>
            </div>
          </CarouselItem>
        ))}
      </CarouselContent>
      <CarouselPrevious className="h-screen" />
      <CarouselNext className="h-[10vh]" />
    </Carousel>
  );
}
