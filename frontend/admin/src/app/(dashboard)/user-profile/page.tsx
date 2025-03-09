"use client";

import { Ellipsis } from "lucide-react";
import Image from "next/image";
import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { Dispatch, SetStateAction, useState } from "react";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { NotificationModal } from "@/components/NotificationModal";

const formSchema = z.object({
  username: z.string().min(2, {
    message: "Username must be at least 2 characters.",
  }),
  fullname: z.string().min(4, {
    message: "Fullname must be at least 4 characters.",
  }),
  email: z.string().email({
    message: "Invalid email address.",
  }),
  country: z.string().min(2, {
    message: "Country must be at least 2 characters.",
  }),
  city: z.string().min(2, {
    message: "City must be at least 2 characters.",
  }),
  town: z.string().min(2, {
    message: "Town must be at least 2 characters.",
  }),
});

export default function UserProfile() {
  const [isOpen, setIsOpen] = useState(false);
  const [isConfirmModal, setIsConfirmModal] = useState(false);
  const [isNotificationModalOpen, setIsNotificationModalOpen] = useState(false);
  const [response, setResponse] = useState<string>("");

  const handleEdit = () => {
    setIsOpen(true);
  };
  return (
    <>
      <div className="flex justify-between mb-10 p-9 border border-[#E6E6E6] rounded-lg">
        <div className="flex gap-6 items-center">
          <div className="relative w-[130px] h-[130px]">
            <Image src="/img/user.jpeg" alt="user" fill className="rounded-full" />
          </div>
          <div className="flex flex-col gap-4 items-start">
            <h3 className="text-2xl font-bold">Ajaladtraveler</h3>
            <p className="text-2xl font-semibold text-[#666666]">Individual</p>
            <span className="bg-celebut-green-light text-celebut-green border-2 rounded-md border-celebut-green px-2 py-1">Active</span>
          </div>
        </div>
        <div className="flex flex-col justify-between">
          <p className="text-[#8C8CA1] text-xl font-normal">
            Account created <span>22/06/2025</span>
          </p>
          <div className="flex items-center justify-end gap-6">
            <Button size="sm" type="button" className="border text-2xl font-normal text-[#999999] flex items-center gap-4 border-[#E6E6E6] rounded-lg bg-white hover:bg-secondary/80 shadow-sm" onClick={handleEdit}>
              Edit{" "}
              <svg width="24" height="25" viewBox="0 0 24 25" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M3 17.7498V21.4998H6.75L17.81 10.4398L14.06 6.68984L3 17.7498ZM21.41 6.83984L17.66 3.08984L15.13 5.62984L18.88 9.37984L21.41 6.83984Z" fill="#666666" />
              </svg>
            </Button>
            <Popover>
              <PopoverTrigger asChild>
                <Button title="more" size="sm" className="border text-2xl border-[#E6E6E6] rounded-lg bg-white hover:bg-secondary/80 shadow-sm">
                  <Ellipsis color="#999999" />
                </Button>
              </PopoverTrigger>
              <PopoverContent className="p-0 w-auto">
                <ul className="flex flex-col">
                  <li className="text-base font-normal">
                    <button type="button" className="hover:bg-secondary/50 w-full text-left py-2 px-4 rounded-md">
                      Warn
                    </button>
                  </li>
                  <li className="text-base font-normal">
                    <button type="button" className="hover:bg-secondary/50 w-full text-left py-2 px-4 rounded-md">
                      Suspend
                    </button>
                  </li>
                  <li className="text-base font-normal">
                    <button type="button" className="hover:bg-secondary/50 w-full text-left py-2 px-4 rounded-md">
                      Deactivate
                    </button>
                  </li>
                  <li className="text-base font-normal">
                    <button type="button" className="hover:bg-secondary/50 w-full text-left py-2 px-4 rounded-md">
                      Restore
                    </button>
                  </li>
                  <li className="text-base font-normal">
                    <button type="button" className="hover:bg-secondary/50 w-full text-left py-2 px-4 rounded-md">
                      Remove
                    </button>
                  </li>
                </ul>
              </PopoverContent>
            </Popover>
          </div>
        </div>
      </div>
      <ul className="flex items-center justify-between p-9 border border-[#E6E6E6] rounded-lg mb-10">
        <li>
          <h3 className="text-lg font-bold text-[#222222]">{(345000).toLocaleString()}</h3>
          <p className="text-[#8C8CA1] text-base font-normal">Following</p>
        </li>
        <li>
          <h3 className="text-lg font-bold text-[#222222]">{(345).toLocaleString()}</h3>
          <p className="text-[#8C8CA1] text-base font-normal">Celebrations</p>
        </li>
        <li>
          <h3 className="text-lg font-bold text-[#222222]">{(345).toLocaleString()}</h3>
          <p className="text-[#8C8CA1] text-base font-normal">Flagged Content</p>
        </li>
        <li>
          <h3 className="text-lg font-bold text-[#222222]">
            {(345).toLocaleString("en-US", {
              style: "currency",
              currency: "USD",
              maximumFractionDigits: 0,
            })}
          </h3>
          <p className="text-[#8C8CA1] text-base font-normal">Cash Gift</p>
        </li>
        <li>
          <h3 className="text-lg font-bold text-[#222222]">
            {(345).toLocaleString("en-US", {
              style: "currency",
              currency: "USD",
              maximumFractionDigits: 0,
            })}
          </h3>
          <p className="text-[#8C8CA1] text-base font-normal">Spent</p>
        </li>
        <li>
          <h3 className="text-lg font-bold text-[#222222]">
            {(345).toLocaleString("en-US", {
              style: "currency",
              currency: "USD",
              maximumFractionDigits: 0,
            })}
          </h3>
          <p className="text-[#8C8CA1] text-base font-normal">Received</p>
        </li>
        <li>
          <h3 className="text-lg font-bold text-[#222222]">
            {(345).toLocaleString("en-US", {
              style: "currency",
              currency: "USD",
              maximumFractionDigits: 0,
            })}
          </h3>
          <p className="text-[#8C8CA1] text-base font-normal">Orders Received</p>
        </li>
      </ul>
      <div className="grid grid-cols-4 gap-9">
        <div className="col-span-3 flex flex-col gap-6 p-9 border border-[#E6E6E6] rounded-lg">
          <div>
            <h3 className="text-[#999999] text-2xl font-semibold">Full Name</h3>
            <p className="text-[#222222] text-2xl font-semibold">Eromosele Oluwatobiloba</p>
          </div>
          <div>
            <h3 className="text-[#999999] text-2xl font-semibold">Email Address</h3>
            <p className="text-[#222222] text-2xl font-semibold">Eromoseleoluwatobiloba@celebut.com</p>
          </div>
          <div>
            <h3 className="text-[#999999] text-2xl font-semibold">Country</h3>
            <p className="text-[#222222] text-2xl font-semibold">Nigeria</p>
          </div>
        </div>
        <div className="col-span-1 p-9 border border-[#E6E6E6] rounded-lg text-[#222222]">
          <h3 className="text-[20px] font-semibold mb-4">Log In Activity</h3>
          <p className="text-base font-normal">Login by</p>
        </div>
      </div>
      <ProfileModal isOpen={isOpen} setIsOpen={setIsOpen} setIsConfirmModal={setIsConfirmModal} />
      <ConfirmProfileModal isOpen={isConfirmModal} setIsConfirmModal={setIsConfirmModal} setIsNotificationModalOpen={setIsNotificationModalOpen} setResponse={setResponse} />
      <NotificationModal
        isOpen={isNotificationModalOpen}
        onClose={() => {
          setIsNotificationModalOpen(false);
        }}
        isSuccess={true}
        response={response}
      />
    </>
  );
}

const ProfileModal = ({ isOpen, setIsOpen, setIsConfirmModal }: { isOpen: boolean; setIsOpen: (isOpen: boolean) => void; setIsConfirmModal: Dispatch<SetStateAction<boolean>> }) => {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      fullname: "",
      username: "",
      email: "",
      country: "",
      city: "",
      town: "",
    },
  });
  function onSubmit(values: z.infer<typeof formSchema>) {
    setIsConfirmModal(true);
    console.log(values);
  }
  return (
    <Dialog open={isOpen} onOpenChange={() => setIsOpen(!isOpen)}>
      <DialogContent className="max-w-2xl max-h-[80%] overflow-y-auto scrollbar-hide">
        <DialogHeader className="w-full">
          <VisuallyHidden>
            <DialogTitle>User Profile</DialogTitle>
          </VisuallyHidden>
        </DialogHeader>
        <div className="p-6 border-t">
          <h5 className="text-base font-bold text-[#8C8CA1] mb-6">Edit Profile</h5>
          <div className="flex flex-col md:flex-row gap-9 mb-7 ">
            <div className="flex flex-col items-center">
              <div className="relative w-[121px] h-[121px] rounded-full mb-4">
                <Image src="/img/user.jpeg" fill className="relative rounded-full" alt="user" />
              </div>
              <div className="flex justify-center items-center">
                <span className="text-xs font-normal bg-celebut-green text-white py-1 px-4 rounded-lg text-center">Active</span>
              </div>
            </div>
            <div className="w-full">
              <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                  <FormField
                    control={form.control}
                    name="fullname"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Full Name</FormLabel>
                        <FormControl>
                          <Input placeholder="John Doe" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="username"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Username</FormLabel>
                        <FormControl>
                          <Input placeholder="johndoe" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="email"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Email</FormLabel>
                        <FormControl>
                          <Input placeholder="johndoe@example.com" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="country"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Country</FormLabel>
                        <FormControl>
                          <Input placeholder="country" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="city"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>City</FormLabel>
                        <FormControl>
                          <Input placeholder="city" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="town"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Town</FormLabel>
                        <FormControl>
                          <Input placeholder="town" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  {/* <Button type="submit">Submit</Button> */}
                  <Button type="submit" className="uppercase bg-celebut-gold text-white w-full text-base font-bold">
                    Submit
                  </Button>
                </form>
              </Form>
            </div>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
};

const ConfirmProfileModal = ({ isOpen, setIsConfirmModal, setIsNotificationModalOpen, setResponse }: { isOpen: boolean; setIsConfirmModal: (isOpen: boolean) => void; setIsNotificationModalOpen: (isOpen: boolean) => void; setResponse: (response: string) => void }) => {
  function handleConfirmation() {
    setIsNotificationModalOpen(true);
    setResponse("Changes Saved Successfully");
    setIsConfirmModal(false);
  }
  return (
    <Dialog open={isOpen} onOpenChange={() => setIsConfirmModal(false)}>
      <DialogContent className="max-w-sm">
        <DialogHeader className="w-full">
          <VisuallyHidden>
            <DialogTitle>Save profile edit</DialogTitle>
          </VisuallyHidden>
        </DialogHeader>
        <div className="flex flex-col mt-6 mb-12 items-center gap-6">
          <h3 className="font-bold text-2xl">Save Changes</h3>
          <p className="font-normal text-base w-[30ch] text-center">Are you sure you want to save the changes made to this user</p>
        </div>
        <DialogFooter>
          <div className="w-full grid grid-cols-1 gap-[18px]">
            <Button type="button" className="uppercase bg-celebut-gold text-white w-full text-base font-bold" onClick={handleConfirmation}>
              Yes, Save
            </Button>
            <Button type="button" className="uppercase text-celebut-gold w-full text-base font-bold bg-transparent hover:bg-transparent" onClick={() => setIsConfirmModal(false)}>
              No, cancel
            </Button>
          </div>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};
