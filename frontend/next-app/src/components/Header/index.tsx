import Image from "next/image"

export default function Header() {
    return(
        <>
            <div className="flex flex-row justify-between p-4 justify-center items-center">
                <Image src="/header/logo.svg" alt="home" width={120} height={120}></Image>
                <div className=" flex flex-row gap-8 mr-4">
                    <div>Home</div>
                    <div className="text-[#FF3F3F]">Register</div>
                    <div>Contact</div>
                    <div>Account</div>
                </div>
            </div>
        </>
    )
};
