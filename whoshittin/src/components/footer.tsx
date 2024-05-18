import {
    Typography,
} from "@material-tailwind/react";

const CURRENT_YEAR = new Date().getFullYear();

export function Footer() {
    return (
        <footer className="pb-5 p-10 md:pt-10">
            <div className="container flex flex-col mx-auto">
                <Typography
                    color="blue-gray"
                    className="text-center mt-12 font-normal !text-gray-700"
                >
                    &copy; {CURRENT_YEAR} Whoshittin{" "}
                    by{" "}
                    <a href="https://zanderambrose.dev" target="_blank">
                        Zander Ambrose
                    </a>
                    .
                </Typography>
            </div>
        </footer>
    );
}

export default Footer;
