import { Footer } from "@/components";
import Hero from "./hero";
import JazzClubs from "./jazzclubs";

import { fetchJazzClubDetails, IVenue } from "@/utils/fetchJazzClubs";

export default async function App() {
    const jazzClubData: Record<string, IVenue[]> = await fetchJazzClubDetails()
    return (
        <>
            <Hero />
            <JazzClubs collections={jazzClubData} />
            <Footer />
        </>
    );
}
