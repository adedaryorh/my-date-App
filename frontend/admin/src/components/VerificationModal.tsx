import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
// import { Label } from "@/components/ui/label";
import { User } from "@/types/interfaces";
import Image from "next/image";
import { Dispatch, SetStateAction } from "react";

export function VerificationModal({ data, isOpen, onClose, setIsBadgeModalOpen, isBadgeModalOpen }: { data: User; isOpen: boolean; setIsBadgeModalOpen: Dispatch<SetStateAction<boolean>>; isBadgeModalOpen: boolean; onClose: () => void }) {
  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="w-auto">
        <DialogHeader className="w-full">
          <VisuallyHidden>
            <DialogTitle>User verification details</DialogTitle>
          </VisuallyHidden>
          <div className="flex items-center gap-6">
            <div className="relative rounded-full flex items-center justify-center w-32 h-32">
              <Image src={data.image} fill alt="user" className="rounded-full" />
            </div>
            <div className="flex flex-col gap-4">
              <h4 className="font-semibold text-base">{data.fullname}</h4>
              <p className="font-normal text-sm">{data.email}</p>
            </div>
          </div>
        </DialogHeader>
        <div className="flex flex-col gap-4 mb-7">
          <div className="grid grid-cols-2">
            <h5 className="font-normal text-sm">Date of birth</h5>
            <p className="font-semibold text-base text-right">{data.dob}</p>
          </div>
          <div className="grid grid-cols-2">
            <h5 className="font-normal text-sm">Location</h5>
            <p className="font-semibold text-base text-right">{`${data.state}, ${data.country}`}</p>
          </div>
          <div className="grid grid-cols-2">
            <h5 className="font-normal text-sm">Application date</h5>
            <p className="font-semibold text-base text-right">{data.application_date}</p>
          </div>
          <div className="grid grid-cols-2">
            <h5 className="font-normal text-sm">Document type</h5>
            <p className="font-semibold text-base text-right">{data.document_type}</p>
          </div>
          <div className="grid grid-cols-2">
            <h5 className="font-normal text-sm">Application status</h5>
            <p className={`${data.status === "pending" ? "text-celebut-gold" : data.status === "approved" ? "text-celebut-green" : "text-celebut-red"} capitalize font-semibold text-base text-right`}>{data.status}</p>
          </div>
        </div>
        <DialogFooter>
          <div className="w-full grid grid-cols-1 gap-[18px]">
            <Button type="button" className="uppercase bg-celebut-gold text-white w-full text-base font-bold" onClick={() => setIsBadgeModalOpen(!isBadgeModalOpen)}>
              Grant verification badge
            </Button>
            {data.status === "approved" && (
              <Button variant={"destructive"} type="button" className="uppercase text-white w-full text-base font-bold">
                Revoke verification badge
              </Button>
            )}
          </div>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
