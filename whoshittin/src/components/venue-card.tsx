import React from "react";
import Image from "next/image";
import {
    Typography,
    Card,
    CardHeader,
    CardBody,
    Avatar,
} from "@material-tailwind/react";
import { IVenue } from "@/utils/fetchJazzClubs";


interface VenueCardProps {
    show: IVenue
}

export function VenueCard({ show }: VenueCardProps) {
    const { eventImage, eventDate, eventLink, eventTime, eventTitle, band, venue } = show
    console.log("eventLink: ", eventLink)
    console.log("eventTime: ", eventTime)
    console.log("band: ", band)
    return (
        <a href={eventLink} target="_blank">
            <Card shadow={true} className="hover:scale-105">
                <CardHeader style={{ margin: "0" }}>
                    <Image
                        width={768}
                        height={768}
                        src={eventImage}
                        alt={venue}
                        className="h-full w-full scale-110 object-cover"
                    />
                </CardHeader>
                <CardBody className="p-6">
                    <Typography
                        variant="h5"
                        color="blue-gray"
                        className="mb-2 normal-case transition-colors hover:text-gray-900"
                    >
                        {eventTitle}
                    </Typography>
                    <div className="container my-auto grid grid-cols-2 gap-4 items-start mb-6">
                        {band?.map(({ name, instrument }) => (
                            <Typography key={`band-member-${name}-${instrument}`} className="font-normal !text-gray-500">
                                {name} - {instrument}
                            </Typography>
                        ))}
                    </div>
                    <div className="flex items-center gap-4">
                        <Avatar
                            size="sm"
                            variant="circular"
                            src={eventImage}
                            alt={eventTitle}
                        />
                        <div>
                            <Typography
                                variant="small"
                                color="blue-gray"
                                className="mb-0.5 !font-medium"
                            >
                                {venue}
                            </Typography>
                            <Typography
                                variant="small"
                                color="gray"
                                className="text-xs !text-gray-500 font-normal"
                            >
                                {eventDate.formattedDate}
                            </Typography>
                        </div>
                    </div>
                </CardBody>
            </Card>
        </a>
    );
}


export default VenueCard;
