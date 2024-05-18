import "./globals.css";
import type { Metadata } from "next";
import { Roboto } from "next/font/google";
import { Layout } from "@/components";

const roboto = Roboto({
    subsets: ["latin"],
    weight: ["300", "400", "500", "700", "900"],
    display: "swap",
});

export const metadata: Metadata = {
    title: "Whoshittin NYC",
    description: "Explore the vibrant jazz scene of New York City with our comprehensive guide to tonight's live performances. Discover top-rated jazz clubs, view event schedules, and find your perfect night out.",
};

export default function RootLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <html lang="en">
            <head>
                {/* TODO - Favicon */}
                <link rel="shortcut icon" href="/favicon.png" type="image/png" />
            </head>
            <body className={roboto.className}>
                <Layout>
                    {children}
                </Layout>
            </body>
        </html>
    );
}
