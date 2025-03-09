/* eslint-disable @next/next/no-img-element */
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import Image from "next/image";

export function NotificationModal({ isOpen, onClose, isSuccess, response }: { isOpen: boolean; isSuccess: boolean; response: string; onClose: () => void }) {
  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="w-auto p-20">
        <DialogHeader className="w-full">
          <VisuallyHidden>
            <DialogTitle>Notification</DialogTitle>
          </VisuallyHidden>
        </DialogHeader>
        <div className="flex flex-col items-center justify-between gap-4 mb-7">
          <div className="relative">{isSuccess ? <Image src="/img/success-icon.png" alt="success" width={100} height={100} /> : <img src={"/img/error-icon.svg"} alt="error" width={100} height={100} />}</div>
        </div>
        <p className="text-[32px] font-semibold text-center">{response}</p>
      </DialogContent>
    </Dialog>
  );
}
