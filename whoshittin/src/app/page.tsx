import { Footer } from "@/components";
import Hero from "./hero";
import JazzClubs from "./jazzclubs";

import { fetchJazzClubs } from "@/utils/fetchJazzClubs";

export default async function App() {
    const jazzClubData: any[] = await fetchJazzClubs()
    return (
        <>
            <Hero />
            <JazzClubs collections={jazzClubData} />
            <Footer />
        </>
    );
}
