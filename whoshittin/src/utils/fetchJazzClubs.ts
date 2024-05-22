export interface IVenue {
    band: {
        instrument: string;
        name: string;
    }[]
    currentTime: string;
    eventDate: {
        date: string;
        formattedDate: string
    }
    eventImage: string
    eventLink: string
    eventTime: {
        start: string
        end: string
    }[]
    eventTitle: string
    venue: string
    _id: string
}

// TODO - This needs env variable
const BASE_API_URL = "http://server:8080/api/v1"
const AUTH_HEADERS = {
    "Authorization": "Bearer admin",
    "X-Client-ID": "admin"
}

const buildTodaysDateQueryParam = () => {
    const today = new Date();
    const options = {
        timeZone: 'America/New_York',
        year: 'numeric',
        month: '2-digit',
        day: '2-digit'
    };

    const formatter = new Intl.DateTimeFormat('en-US', options as Intl.DateTimeFormatOptions);
    const [{ value: month }, , { value: day }, , { value: year }] = formatter.formatToParts(today);

    const formattedDate = `${year}-${month}-${day}`;
    console.log("formattedDate:", formattedDate);

    return formattedDate;
}

export async function fetchJazzClubCollections() {
    const res = await fetch(`${BASE_API_URL}/collections`, {
        headers: {
            ...AUTH_HEADERS
        }
    })

    const { collections } = await res.json()

    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return collections
}

export async function fetchJazzClubDetails() {
    const dateString = buildTodaysDateQueryParam()
    const collections: string[] = await fetchJazzClubCollections()
    const fetchVenueDetails: Promise<Response>[] = []
    const jazzClubsMap: Record<string, IVenue[]> = {}

    for (const collection of collections) {
        fetchVenueDetails.push(fetch(`${BASE_API_URL}/${collection}?date=${dateString}`, {
            headers: {
                ...AUTH_HEADERS
            }
        }))
    }

    const res = await Promise.all(fetchVenueDetails)

    for (const response of res) {
        const data: IVenue[] = await response.json()
        if (data) {
            jazzClubsMap[data[0].venue] = data
        }
    }

    return jazzClubsMap
}
