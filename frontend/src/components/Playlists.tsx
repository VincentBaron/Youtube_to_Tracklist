import * as React from "react";

import { Card, CardContent } from "@/components/ui/card";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";

export function Playlists() {
  return (
    <div>
      <Carousel
        opts={{
          align: "center",
        }}
        orientation="horizontal"
      >
        <CarouselContent>
          {Array.from({ length: 3 }).map((_, index) => (
            <CarouselItem key={index}>
              <span>yolo</span>
            </CarouselItem>
          ))}
        </CarouselContent>
        <CarouselPrevious />
        <CarouselNext />
      </Carousel>
    </div>
  );
}
