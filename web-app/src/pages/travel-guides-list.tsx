import DefaultLayout from "@/layouts/default";
import {
    Category,
    createTravelGuide,
    CreateTravelGuide,
} from "@/service/TravelGuide";
import { Button } from "@nextui-org/button";
import { Checkbox } from "@nextui-org/checkbox";
import { Divider } from "@nextui-org/divider";
import { Form } from "@nextui-org/form";
import { Input } from "@nextui-org/input";
import {
    Modal,
    ModalBody,
    ModalContent,
    ModalHeader,
    useDisclosure,
} from "@nextui-org/modal";
import { Popover, PopoverContent, PopoverTrigger } from "@nextui-org/popover";
import { Select, SelectItem } from "@nextui-org/select";

export default function TravelGuidesListPage() {
    const { isOpen, onOpen, onOpenChange, onClose } = useDisclosure();

    const onSubmit = async (e: any) => {
        e.preventDefault();
        const data = Object.fromEntries(
            new FormData(e.currentTarget),
        ) as Record<string, string>;

        const createTravelGuideRequest: CreateTravelGuide = {
            travelGuide: {
                name: data.name,
                description: data.description,
                isPrivate: data.isPrivate === undefined ? false : true,
                location: {
                    street: data.street,
                    zip: data.zip,
                    city: data.city,
                    state: data.state,
                    country: data.country,
                },
                category: parseInt(
                    data.category ?? Category.MIX,
                ) as unknown as Category,
            },
            secret: data.password ?? "",
        };
        try {
            await createTravelGuide(createTravelGuideRequest);
            onClose();
        } catch (e) {}
    };

    const categories = [
        { key: 0, label: "Mixed" },
        { key: 1, label: "Culture" },
        { key: 2, label: "Action" },
        { key: 3, label: "Relax" },
        { key: 4, label: "Adventure" },
        { key: 5, label: "Sports" },
        { key: 6, label: "Roadtrip" },
    ];

    return (
        <DefaultLayout>
            <section>
                <h2 className="text-3xl text-primary-500">Travel Guides</h2>
                <p className="mt-3">
                    Here you can find a list of Travel Guides or create your
                    own.
                </p>
                <Button className="mt-3" color="primary" onPress={onOpen}>
                    Create Travel Guide
                </Button>

                {/* <div className="mt-4">
                    <Alert
                        key="created-tg-alert"
                        color="success"
                        title={`Created Travel Guide`}
                        description=""
                        variant="flat"
                    />
                    <Alert
                        key="error-tg-alert"
                        color="danger"
                        title={`Creation of Travel Guide failed.`}
                        description=""
                        variant="flat"
                    />
                </div> */}

                {/* Create Travel Guide Modal */}
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
                                    Create new Travel Guide
                                </ModalHeader>
                                <ModalBody>
                                    <Form
                                        className="grid"
                                        validationBehavior="native"
                                        onSubmit={onSubmit}
                                    >
                                        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                                            <div>
                                                <Input
                                                    name="name"
                                                    className="mt-2"
                                                    label="Name"
                                                    placeholder="Name your Travel Guide"
                                                    isRequired
                                                    isClearable
                                                    minLength={1}
                                                />
                                                <Input
                                                    name="description"
                                                    className="mt-2"
                                                    label="Description"
                                                    placeholder="Details about your Travel Guide"
                                                    isClearable
                                                />
                                                <Select
                                                    name="category"
                                                    className="mt-2"
                                                    label="Category"
                                                >
                                                    {categories.map(
                                                        (category) => (
                                                            <SelectItem
                                                                key={
                                                                    category.key
                                                                }
                                                            >
                                                                {category.label}
                                                            </SelectItem>
                                                        ),
                                                    )}
                                                </Select>
                                            </div>
                                            <div>
                                                <Input
                                                    name="street"
                                                    className="mt-2"
                                                    label="Street"
                                                    placeholder="Street"
                                                    isClearable
                                                />
                                                <Input
                                                    name="zip"
                                                    className="mt-2"
                                                    label="Zip Code"
                                                    placeholder="Zip Code"
                                                    isClearable
                                                />
                                                <Input
                                                    name="city"
                                                    className="mt-2"
                                                    label="City"
                                                    placeholder="City"
                                                    isClearable
                                                />
                                                <Input
                                                    name="state"
                                                    className="mt-2"
                                                    label="State"
                                                    placeholder="State"
                                                    isClearable
                                                />
                                                <Input
                                                    name="country"
                                                    className="mt-2"
                                                    label="Country"
                                                    placeholder="Country"
                                                    isRequired
                                                    isClearable
                                                />

                                                <Popover>
                                                    <PopoverTrigger>
                                                        <div>
                                                            <Input
                                                                className="mt-2"
                                                                label="Planet"
                                                                placeholder=""
                                                                value={
                                                                    "🌎 Earth"
                                                                }
                                                            />
                                                        </div>
                                                    </PopoverTrigger>
                                                    <PopoverContent>
                                                        <div className="px-1 py-2">
                                                            <div className="text-small">
                                                                Sorry, currently
                                                                you can only
                                                                travel to places
                                                                on planet earth.
                                                            </div>
                                                            <div className="text-tiny">
                                                                But good to know
                                                                you're
                                                                interested in
                                                                visisting other
                                                                planets. 😅
                                                            </div>
                                                        </div>
                                                    </PopoverContent>
                                                </Popover>
                                            </div>
                                        </div>

                                        <Divider className="mt-2"></Divider>

                                        <div>
                                            <div className="flex py-2 px-1 justify-between">
                                                <Checkbox
                                                    name="isPrivate"
                                                    size="md"
                                                    classNames={{
                                                        label: "text-small",
                                                    }}
                                                >
                                                    Private
                                                    <br />
                                                    <span className="text-gray-500">
                                                        A private Travel Guide
                                                        can only be viewed with
                                                        Password. The name is
                                                        always public.
                                                    </span>
                                                </Checkbox>
                                            </div>

                                            <Input
                                                name="password"
                                                className="mt-2"
                                                label="Password"
                                                placeholder="Enter your password"
                                                type="password"
                                                isRequired
                                                minLength={8}
                                                description="The password is required to edit your Travel Guides and to view it when it is private."
                                            />
                                        </div>
                                        <div className="flex justify-end mb-4">
                                            <div className="grow-1"></div>
                                            <Button
                                                color="danger"
                                                onPress={onClose}
                                                className="mr-2"
                                            >
                                                Cancel
                                            </Button>
                                            <Button
                                                color="primary"
                                                type="submit"
                                            >
                                                Create Travel Guide
                                            </Button>
                                        </div>
                                    </Form>
                                </ModalBody>
                            </>
                        )}
                    </ModalContent>
                </Modal>
            </section>
        </DefaultLayout>
    );
}
