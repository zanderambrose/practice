"use client";

import React from "react";
import {
    Tabs,
    TabsHeader,
    Tab,
} from "@material-tailwind/react";
import VenueCard from "@/components/venue-card";

const venues = []

export function Posts() {
    return (
        <section className="grid min-h-screen place-items-center p-8">
            <Tabs value="jazz" className="mx-auto max-w-7xl w-full mb-16 ">
                <div className="w-full flex mb-8 flex-col items-center">
                    <TabsHeader className="h-10 !w-12/12 md:w-[50rem] border border-white/25 bg-opacity-90">
                        <Tab value="jazz">Jazz</Tab>
                        <Tab value="comedy">Comedy</Tab>
                        <Tab value="theater">Theater</Tab>
                    </TabsHeader>
                </div>
            </Tabs>
            <div className="container my-auto grid grid-cols-1 gap-x-8 gap-y-16 items-start lg:grid-cols-3">
                {venues.map(({ img, tag, title, desc, date, author }) => (
                    <VenueCard
                        key={title}
                        img={img}
                        tag={tag}
                        title={title}
                        desc={desc}
                        date={date}
                        author={{
                            img: author.img,
                            name: author.name,
                        }}
                    />
                ))}
            </div>
        </section>
    );
}

export default Posts;
