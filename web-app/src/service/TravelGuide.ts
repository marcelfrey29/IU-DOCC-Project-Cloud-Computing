export interface TravelGuide {
    id?: string;
    name: string;
    description?: string;
    isPrivate: boolean;
    location: Location;
    category: Category;
}

export interface Location {
    street?: string;
    zip?: string;
    city?: string;
    state?: string;
    country: string;
}

export enum Category {
    MIX,
    CULTURE,
    ACTION,
    RELAX,
    ADVENTURE,
    SPORTS,
    ROADTRIP,
}

export interface CreateTravelGuide {
    travelGuide: TravelGuide;
    secret: string;
}

export async function createTravelGuide(
    data: CreateTravelGuide,
): Promise<TravelGuide> {
    const response = await fetch("http://localhost:9090/travel-guides", {
        method: "POST",
        headers: new Headers({
            "content-type": "application/json",
        }),
        body: JSON.stringify(data),
    });
    if (response.status !== 201) {
        throw new Error();
    }
    const body = await response.json();
    return body;
}
