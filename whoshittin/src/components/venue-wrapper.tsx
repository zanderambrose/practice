import React from "react";
import { IVenue } from "@/utils/fetchJazzClubs";
import VenueCard from "@/components/venue-card";
import { Typography } from "@material-tailwind/react";
import { venueMap } from "@/utils/venueMap";

interface VenueWrapperProps {
    shows: IVenue[];
    venueName: string;
}

export function VenueWrapper({ shows, venueName }: VenueWrapperProps) {
    return (
        <>
            <Typography variant="h1" className="mb-4">{venueMap[venueName]}</Typography>
            <div className="container my-auto grid grid-cols-1 gap-x-16 gap-y-16 items-start lg:grid-cols-2">
                {shows.map((show) => (
                    <VenueCard key={show._id} show={show} />
                ))}
            </div>
        </>
    );
}

export default VenueWrapper;
