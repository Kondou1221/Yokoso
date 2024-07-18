import Image from "next/image"
export default function Register() {
    return(
        <>
            <div className="flex flex-row gap-8 p-4 w-full h-full">
                <div className="flex flex-col gap-2 w-[60%] ml-8">
                    <div className="text-[#FF3F3F] text-2xl">REGISTRATION</div>
                    <div className="font-bold text-8xl">Register</div>
                    <div className="font-bold text-8xl">Map</div>
                    <div className="ml-12 mt-[-15%]">
                        <Image src="/register/iPhone.svg" alt="home" width={550} height={550} ></Image>
                    </div>
                </div>
                <form action="" className="flex-col rounded-lg p-4 w-[40%] shadow-xl shadow-yl mr-8 h-1/2">
                    <div className="flex flex-col justify-center items-center gap-4 w-full overflow-y-auto">
                        <div className="flex flex-col justify-center items-center gap-1">
                            <Image src="/register/location.svg" alt="home" width={200} height={200}></Image>
                            <div className="text-[#FF3F3F] font-bold text-base">Place Registration</div>
                            <div>Fill out the form with your information</div>
                        </div>
                        {/* 店舗 */}
                        <div className="flex flex-col justify-center items-center gap-2 w-full">
                            <div className="flex flex-row gap-2 w-full">
                                <input type="text" placeholder="店舗名" className="border border-[#D9D9D9] rounded-md pl-2 w-[70%]"/>
                                <select name="subcategory" id="" className="border border-[#D9D9D9] rounded-md p-2 w-[30%]">
                                    <option value="food">Food</option>
                                    <option value="attraction">attraction</option>
                                    <option value="shop">Shop</option>
                                </select>
                            </div>
                            <textarea name="description" id="" placeholder="店舗の説明" cols={20} rows={7} className="border border-[#D9D9D9] rounded-md pl-2 w-full"></textarea>
                        </div>
                        {/* 住所 */}
                        <input type="text" placeholder="住所" className="border border-[#D9D9D9] rounded-md w-full p-1"/>
                        {/* 価格 */}
                        <div className="flex flex-row gap-2 w-full justify-start items-center">
                            <input type="number" placeholder="最安値" className="border border-[#D9D9D9] rounded-md p-1 pl-3 w-[30%]"/>
                            <div>~</div>
                            <input type="number" name="" id="" placeholder="最高値"  className="border border-[#D9D9D9] rounded-md p-1 pl-3 w-[30%]"/>
                        </div>
                        {/* 曜日 */}
                        <div className="flex flex-col justify-center gap-2 w-full">
                            <div className="flex flex-row gap-8 items-center">
                                <div>月曜日</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                                <div>~</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                            </div>
                            <div className="flex flex-row gap-8 items-center">
                                <div>火曜日</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                                <div>~</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                            </div>
                            <div className="flex flex-row gap-8 items-center">
                                <div>水曜日</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                                <div>~</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                            </div>
                            <div className="flex flex-row gap-8 items-center">
                                <div>木曜日</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                                <div>~</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                            </div>
                            <div className="flex flex-row gap-8 items-center">
                                <div>金曜日</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                                <div>~</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                            </div>
                            <div className="flex flex-row gap-8 items-center">
                                <div>土曜日</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                                <div>~</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                            </div>
                            <div className="flex flex-row gap-8 items-center">
                                <div>日曜日</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                                <div>~</div>
                                <input type="time" className="border border-[#D9D9D9] rounded-md text-center" />
                            </div>
                        </div>
                        {/* タグ */}
                        <div className="flex flex-row gap-2 w-full">
                            <input type="text" placeholder="タグ入力" className="border border-[#D9D9D9] rounded-md p-1 w-[70%]"/>
                            <button className="bg-[#D9D9D9] rounded-md p-1 w-[20%]">追加</button>
                        </div>
                        <div className="flex flex-row gap-2 items-start w-full">
                            <div className="bg-[#D9D9D9] rounded-md p-1">タグ</div>
                            <div className="bg-[#D9D9D9] rounded-md p-1">タグ</div>
                        </div>
                        {/* 画像 */}
                        <div className="flex flex-row justify-start items-center gap-2 w-[60%] border border-[#434343] rounded-md p-2">
                            <Image src="/register/image.svg" alt="home" width={35} height={35}></Image>
                            <input type="file" className="opacity-0 w-full"/>
                            <div className="ml-[-90%]">画像を追加</div>
                        </div>
                        {/* 登録ボタン */}
                        <div className="flex flex-col w-full justify-center items-center">
                            <button className="bg-[#FF3F3F] text-[#ffffff] w-[90%] rounded-md p-2">Register</button>
                        </div>
                    </div>
                </form>
            </div>
        </>
    )
};
