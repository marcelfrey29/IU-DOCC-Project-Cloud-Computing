import { ActivityEditor } from "@/components/activity-editor";
import { BootstrapIcon } from "@/components/icons";
import { TravelGuideEditor } from "@/components/travel-guide-editor";
import { categoryConfig } from "@/config/category";
import DefaultLayout from "@/layouts/default";
import {
    deleteTravelGuideById,
    getTravelGuideById,
    TravelGuide,
} from "@/service/TravelGuide";
import { Alert } from "@nextui-org/alert";
import { Button } from "@nextui-org/button";
import { Input } from "@nextui-org/input";
import { Link } from "@nextui-org/link";
import {
    Modal,
    ModalBody,
    ModalContent,
    ModalHeader,
    useDisclosure,
} from "@nextui-org/modal";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

export default function TravelGuideDetailPage() {
    let { id } = useParams();
    const routeTo = useNavigate();
    const { isOpen, onOpen, onOpenChange, onClose } = useDisclosure();
    const [travelGuide, setTravelGuide] = useState(null as TravelGuide | null);
    const [authRequited, setAuthRequired] = useState(true);
    const [deleteError, setDeleteError] = useState(false);
    const [secret, setSecret] = useState("");

    const getTravelGuideData = async () => {
        try {
            setTravelGuide(await getTravelGuideById(id ?? "unknown", secret));
            setAuthRequired(false);
        } catch (e) {
            setAuthRequired(true);
            setSecret("");
        }
    };
    useEffect(() => {
        getTravelGuideData();
    }, []);

    const deleteTravelGuide = async () => {
        try {
            await deleteTravelGuideById(id ?? "unknown", secret);
            setDeleteError(false);
            onClose();
            routeTo("/travel-guides");
        } catch (e) {
            setDeleteError(true);
        }
        setSecret("");
    };

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
                                    <p className="">
                                        <BootstrapIcon
                                            name="tag-fill"
                                            className="mr-1"
                                        ></BootstrapIcon>
                                        {
                                            categoryConfig.find(
                                                (c) =>
                                                    c.key ===
                                                    travelGuide.category,
                                            )?.label
                                        }
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

                                    <div className="flex">
                                        {/* Edit Action */}
                                        <div className="mr-2">
                                            <TravelGuideEditor
                                                type="update"
                                                data={travelGuide}
                                                onSuccess={(d) => {
                                                    setTravelGuide(d);
                                                }}
                                            ></TravelGuideEditor>
                                            {/*TODO: Add Edit Travel Guide */}
                                        </div>

                                        {/* Delete Action */}
                                        <div>
                                            <Button
                                                className="mt-3"
                                                color="danger"
                                                onPress={onOpen}
                                            >
                                                <BootstrapIcon name="trash-fill"></BootstrapIcon>
                                                Create Travel Guide
                                            </Button>

                                            {/* Delete Travel Guide Modal */}
                                            <Modal
                                                isOpen={isOpen}
                                                onOpenChange={onOpenChange}
                                                backdrop="blur"
                                                size="4xl"
                                                scrollBehavior="inside"
                                            >
                                                <ModalContent>
                                                    {(onClose) => (
                                                        <>
                                                            <ModalHeader className="flex flex-col gap-1">
                                                                Delete Travel
                                                                Guide
                                                            </ModalHeader>
                                                            <ModalBody>
                                                                <p>
                                                                    Are you sure
                                                                    that you
                                                                    want to
                                                                    delete this
                                                                    Travel
                                                                    Guide?
                                                                </p>
                                                                <p>
                                                                    To delete
                                                                    the Travel
                                                                    Guide, the
                                                                    password is
                                                                    required.
                                                                </p>
                                                                {deleteError ? (
                                                                    <>
                                                                        <Alert
                                                                            description="Error while deleting the Travel Guide"
                                                                            color="danger"
                                                                        >
                                                                            Check
                                                                            the
                                                                            password
                                                                            and
                                                                            try
                                                                            again.
                                                                        </Alert>
                                                                    </>
                                                                ) : (
                                                                    <></>
                                                                )}
                                                                <div>
                                                                    <Input
                                                                        name="password"
                                                                        className="mt-3"
                                                                        label="Password"
                                                                        placeholder="Travel Guide Password"
                                                                        type="password"
                                                                        isRequired
                                                                        value={
                                                                            secret
                                                                        }
                                                                        onValueChange={
                                                                            setSecret
                                                                        }
                                                                    />
                                                                </div>
                                                                <div className="flex justify-end mb-4 mt-3">
                                                                    <div className="grow-1"></div>
                                                                    <Button
                                                                        color="default"
                                                                        onPress={
                                                                            onClose
                                                                        }
                                                                        className="mr-2"
                                                                    >
                                                                        <BootstrapIcon name="x-circle-fill"></BootstrapIcon>
                                                                        Cancel
                                                                    </Button>
                                                                    <Button
                                                                        color="danger"
                                                                        onPress={
                                                                            deleteTravelGuide
                                                                        }
                                                                        isDisabled={
                                                                            secret.length <
                                                                            1
                                                                        }
                                                                    >
                                                                        <BootstrapIcon name="trash-fill"></BootstrapIcon>
                                                                        Delete
                                                                    </Button>
                                                                </div>
                                                            </ModalBody>
                                                        </>
                                                    )}
                                                </ModalContent>
                                            </Modal>
                                        </div>
                                    </div>
                                </section>

                                {/* Activities */}
                                <section className="mt-10">
                                    <h2 className="text-2xl">
                                        <b>Activities</b>
                                    </h2>
                                    <ActivityEditor
                                        type="create"
                                        travelGuideId={id ?? "unknown"}
                                        onSuccess={() => {
                                            /* TODO: add list update */
                                        }}
                                    ></ActivityEditor>
                                </section>
                            </section>
                        </>
                    )}
                </>
            )}
        </DefaultLayout>
    );
}
