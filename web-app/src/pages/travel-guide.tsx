import { BootstrapIcon } from "@/components/icons";
import DefaultLayout from "@/layouts/default";
import { getTravelGuideById, TravelGuide } from "@/service/TravelGuide";
import { Button } from "@nextui-org/button";
import { Input } from "@nextui-org/input";
import { Link } from "@nextui-org/link";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

export default function TravelGuideDetailPage() {
    let { id } = useParams();
    const routeTo = useNavigate();
    const [travelGuide, setTravelGuide] = useState(null as TravelGuide | null);
    const [authRequited, setAuthRequired] = useState(true);
    const [secret, setSecret] = useState("");

    const getTravelGuideData = async () => {
        try {
            setTravelGuide(await getTravelGuideById(id ?? "unknown", secret));
            setAuthRequired(false);
        } catch (e) {
            setAuthRequired(true);
        }
        setSecret("");
    };
    useEffect(() => {
        getTravelGuideData();
    }, []);

    return (
        <DefaultLayout>
            <section className="mb-4">
                <Link
                    href="/travel-guides"
                    color="primary"
                    className="text-sm hover:underline"
                >
                    <BootstrapIcon name="chevron-left" className="mr-1" /> All
                    Travel Guides
                </Link>
            </section>

            {authRequited === true ? (
                <>
                    {/* Auth */}
                    <section className="mt-6 mb-12">
                        <p className="font-mono text-3xl">
                            <BootstrapIcon
                                name="lock-fill"
                                className="text-danger"
                            ></BootstrapIcon>{" "}
                            <b>401</b>
                        </p>
                        <p className="text-3xl">
                            This Travel Guide is Private and requires a
                            Password.
                        </p>

                        <Input
                            name="secret"
                            className="mt-8"
                            label="Password"
                            type="password"
                            placeholder="Password of the Travel Guide"
                            isRequired
                            isClearable
                            minLength={1}
                            value={secret}
                            onValueChange={setSecret}
                        />

                        <Button
                            className="mt-4"
                            color="primary"
                            onPress={getTravelGuideData}
                        >
                            <BootstrapIcon name="unlock-fill" /> Unlock
                        </Button>
                    </section>
                </>
            ) : (
                <>
                    {travelGuide === null ? (
                        <>
                            {/* Not Found */}
                            <section className="mt-6 mb-12">
                                <p className="font-mono text-3xl">
                                    <BootstrapIcon
                                        name="signpost-split"
                                        className="text-danger"
                                    ></BootstrapIcon>{" "}
                                    <b>404</b>
                                </p>
                                <p className="text-3xl">
                                    This Travel Guide doesn't exist.
                                </p>

                                <Button
                                    className="mt-4"
                                    onPress={() => routeTo(`/travel-guides`)}
                                    color="primary"
                                >
                                    <BootstrapIcon name="list-task" /> View all
                                    Travel Guides
                                </Button>
                            </section>
                        </>
                    ) : (
                        <>
                            {/* Travel Guide View */}
                            <section>
                                {/* Title */}
                                <section>
                                    <h2 className="text-3xl text-primary-500">
                                        <b>{travelGuide.name}</b>
                                    </h2>
                                    <p className="mt-1 text-xl">
                                        {travelGuide.description}
                                    </p>
                                    <p className="mt-3 ">
                                        <BootstrapIcon
                                            name="geo-alt-fill"
                                            className="mr-1"
                                        ></BootstrapIcon>
                                        {[
                                            travelGuide.location.street,
                                            `${travelGuide.location.zip} ${travelGuide.location.city}`.trim(),
                                            travelGuide.location.state,
                                            travelGuide.location.country,
                                        ]
                                            .filter(
                                                (e) =>
                                                    e !== undefined && e !== "",
                                            )
                                            .join(" ▪ ")}
                                    </p>
                                    <p>
                                        {travelGuide.isPrivate === true ? (
                                            <>
                                                <BootstrapIcon
                                                    name="eye-slash-fill"
                                                    className="mr-1"
                                                ></BootstrapIcon>
                                                <span>Private</span>
                                            </>
                                        ) : (
                                            <>
                                                <BootstrapIcon
                                                    name="eye-fill"
                                                    className="mr-1"
                                                ></BootstrapIcon>
                                                <span>Public</span>
                                            </>
                                        )}
                                    </p>
                                    {/*TODO: Add Edit/Delete Travel Guide */}
                                </section>

                                {/* Activities */}
                                <section className="mt-10">
                                    <h2 className="text-2xl">
                                        <b>Activities</b>
                                    </h2>
                                    <p className="mt-1">No Activies.</p>
                                    {/* TODO: Add Activities & Activity Management */}
                                </section>
                            </section>
                        </>
                    )}
                </>
            )}
        </DefaultLayout>
    );
}
