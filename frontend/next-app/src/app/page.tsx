import Register from "@/components/Register";
import Header from "@/components/Header";
export default function Home() {
  return (
    <div className="w-full h-full flex flex-col">
      <Header/>
      <Register/>
    </div>
  );
}
