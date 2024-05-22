"use client";

import React from "react";
import {
    Tabs,
    TabsHeader,
    Tab,
} from "@material-tailwind/react";
import VenueWrapper from "@/components/venue-wrapper";
import { IVenue } from "@/utils/fetchJazzClubs";

interface IJazzClubsProps {
    collections: Record<string, IVenue[]>
}

export function JazzClubs({ collections }: IJazzClubsProps) {
    return (
        <section className="grid min-h-screen place-items-center p-8">
            <Tabs value="jazz" className="mx-auto max-w-7xl w-full mb-16 ">
                <div className="w-full flex mb-8 flex-col items-center">
                    <TabsHeader className="h-10 !w-12/12 md:w-[50rem] border border-white/25 bg-opacity-90">
                        <Tab value="jazz">Jazz</Tab>
                        <Tab value="comedy">Comedy</Tab>
                    </TabsHeader>
                </div>
            </Tabs>
            <div className="container flex flex-col gap-8">
                {Object.values(collections).map((venue) => (
                    <VenueWrapper key={venue[0]._id} shows={venue} venueName={venue[0].venue} />
                ))}
            </div>
        </section>
    );
}

export default JazzClubs;
