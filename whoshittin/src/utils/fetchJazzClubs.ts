// TODO - This needs env variable
const BASE_API_URL = "http://server:8080/api/v1"
const AUTH_HEADERS = {
    "Authorization": "Bearer admin",
    "X-Client-ID": "admin"
}

export async function fetchJazzClubs() {
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

// export async function fetchJazzClubDetail() {
//     const collections = await fetchJazzClubs()
// }
