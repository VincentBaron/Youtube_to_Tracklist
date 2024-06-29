import * as React from "react";
import { useState, useEffect } from "react";
import axios, { AxiosResponse } from "axios";

import { Card, CardContent } from "@/components/ui/card";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";

export default function SetsPage() {
  const [sets, setSets] = useState<{ link: string }[]>([]);

  useEffect(() => {
    axios
      .get("http://localhost:8080/sets", { withCredentials: true })
      .then((response: AxiosResponse<any>) => {
        console.log(response.data);
        setSets(response.data.sets);
      });
  }, []);

  return (
    <div>
      <Carousel className="h-[20vh] w-[20vw]">
        <CarouselContent>
          {sets.map((set, index) => (
            <CarouselItem key={index}>
              <Card>
                <CardContent className="flex aspect-square items-center justify-center">
                  <span className="text-4xl font-semi bold">{set.link}</span>
                </CardContent>
              </Card>
            </CarouselItem>
          ))}
        </CarouselContent>
        <CarouselPrevious />
        <CarouselNext />
      </Carousel>
    </div>
  );
}
