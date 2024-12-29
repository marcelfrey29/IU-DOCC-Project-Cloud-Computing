import { Category, Location } from "./Shared";

/**
 * An Activity is one element of a Travel Guide.
 */
export interface Activity {
    id?: string;
    name: string;
    description?: string;
    location: Location;
    category: Category;
    timeInMin: number;
    costsInCent: number;
}

export interface CreateActivityRequest {
    activity: Activity;
}

export async function createActivity(
    travelGuideId: string,
    data: CreateActivityRequest,
    secret: string,
): Promise<Activity[]> {
    const response = await fetch(
        "http://localhost:9090/travel-guides/" + travelGuideId + "/activities",
        {
            method: "POST",
            headers: new Headers({
                "content-type": "application/json",
                "x-tg-secret": secret,
            }),
            body: JSON.stringify(data),
        },
    );
    if (response.status !== 201) {
        throw new Error();
    }
    const body = await response.json();
    return body;
}
