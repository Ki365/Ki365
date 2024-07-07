import { PATH } from "@/routes/routes";
import { BookOpenIcon, CpuIcon, FoldersIcon, PartyPopperIcon, ShieldEllipsisIcon, UserRoundCogIcon, UsersIcon } from "lucide-react";
import { Link, useLocation } from "react-router-dom";
import { Separator } from "@/components/ui/separator";
import { ReactNode } from "react";

interface SidebarProps {
	/**
	 * @options `true/false`
	 * @description If false, unhides a space component. Useful to differentiate mobile and desktop views.
	 */
	adaptive: boolean
}

interface ButtonLinkProps {
	/**
	 * @options `any PATH string or blank`
	 * @description Required to link to a page and show button as "selected." Leave blank to point to no link.
	 */
	path: string
	/**
	 * @options `any string, should start with capital letter`
	 * @description Text to show inside button.
	 */
	text: string
	/**
	 * @options `any ReactNode component, should be icon component`
	 * @description Logo to show inside button.
	 */
	children: ReactNode
}

const buttonLinkClass = (
	"flex items-center gap-3 rounded-lg px-3 py-2 trainsition-all hover:text-primary text-muted-foreground"
)

const iconClass = (
	"h-4 w-4"
)

function ButtonLink(props: ButtonLinkProps) {
	const {pathname} = useLocation()

	return (
		<Link to={props.path}
			className={`${buttonLinkClass} ${pathname.startsWith(props.path) ? 'text-primary bg-muted' : 'text-muted-foreground'} `}>
			{props.children}
			{/* <BookOpenIcon className="h-4 w-4" /> */}
			{props.text}
		</Link>
	)
}

// TODO: render this component based on input props (to enable different users)
export default function Sidebar(props : SidebarProps) {
	return (
		<nav className="grid items-start px-2 text-sm font-medium lg:px-4">
			<br className={`${props.adaptive ? 'hidden' : ''}`} />

			<ButtonLink path={PATH.USER.HOME} text="Home">
				<BookOpenIcon className={iconClass} />
			</ButtonLink>
			<Separator className="my-2" />

			<ButtonLink path={PATH.USER.ALL_PROJECTS} text="Projects">
				<FoldersIcon className="h-4 w-4" />
			</ButtonLink>
			<ButtonLink path={PATH.USER.COMPONENTS} text="Components">
				<CpuIcon className="h-4 w-4" />
			</ButtonLink>

			<Separator className="my-2" />

			<ButtonLink path={PATH.USER.MEMBERS} text="Members">
				<UsersIcon className="h-4 w-4" />
			</ButtonLink>
			<ButtonLink path={PATH.USER.PREFERENCES} text="Preferences">
				<UserRoundCogIcon className="h-4 w-4" />
			</ButtonLink>

			<Separator className="my-2" />

			<ButtonLink path={PATH.ADMIN.ADMIN_SETTINGS} text="Admin">
				<ShieldEllipsisIcon className="h-4 w-4" />
			</ButtonLink>

			<Separator className="my-2" />

			<br className="" />
			<br className="" />

			<ButtonLink path={PATH.USER.SURPRISE} text="Surprise Me!">
				<PartyPopperIcon className="h-4 w-4" />
			</ButtonLink>

		</nav>
	)
}