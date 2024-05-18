"use client";

import Image from "next/image";
import { Typography } from "@material-tailwind/react";


function Hero() {
    return (
        <header className="mt-5 bg-white p-8">
            <div className="w-w-full container mx-auto p-12 text-center">
                <Typography
                    color="blue-gray"
                    className="mx-auto w-full text-[30px] lg:text-[48px] font-bold leading-[45px] lg:leading-[60px] lg:max-w-2xl"
                >
                    Whoshittin NYC!
                </Typography>
                <Typography
                    variant="lead"
                    className="mx-auto mt-8 mb-4 w-full px-8 !text-gray-700 lg:w-10/12 lg:px-12 xl:w-8/12 xl:px-20"
                >
                    New York Cities Premier Jazz Club Explorer
                </Typography>
            </div>
            <div className="w-full lg:container lg:mx-auto">
                <Image
                    width={1024}
                    height={400}
                    src="/image/whoshittin-jazz-hero.png"
                    alt="credit cards"
                    className="h-96 w-full rounded-lg object-cover lg:h-[21rem]"
                />
            </div>
        </header>
    );
}
export default Hero;
