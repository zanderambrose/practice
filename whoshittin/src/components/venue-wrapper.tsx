import React from "react";
import { IVenue } from "@/utils/fetchJazzClubs";
import VenueCard from "@/components/venue-card";

interface VenueWrapperProps {
    shows: IVenue[]
}

export function VenueWrapper({ shows }: VenueWrapperProps) {
    return (
        <div className="container my-auto grid grid-cols-1 gap-x-8 gap-y-16 items-start lg:grid-cols-3">
            {shows.map((show) => (
                <VenueCard key={show._id} show={show} />
            ))}
        </div>
    );
}


export default VenueWrapper;
