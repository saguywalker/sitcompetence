import { useState, useEffect, useRef, useCallback } from "react";

export function useHover() {
	const [isHover, setIsHover] = useState(false);
	const element = useRef(null);

	const setElementRef = useCallback((el) => element.current = el, []);

	const onMouseEnter = useCallback(() => setIsHover(true), []);
	const onMouseLeave = useCallback(() => setIsHover(false), []);

	useEffect(() => {
		if (element.current !== null) {
			element.current.addEventListener("mouseenter", onMouseEnter);
			element.current.addEventListener("mouseleave", onMouseLeave);
		}

		return () => {
			if (element.current !== null) {
				element.current.removeEventListener("mouseenter", onMouseEnter);
				element.current.removeEventListener("mouseleave", onMouseLeave);
			}
		};
	}, [onMouseEnter, onMouseLeave]);

	return [setElementRef, isHover];
}

export function useClickOutside() {
	const [clickOutside, setClickOutside] = useState(true);
	const element = useRef(null);

	const setElementRef = useCallback((el) => element.current = el, []);

	const onClickOutside = useCallback(() => {
		if (element.current) {
			setClickOutside(false);
		}
	}, []);

	useEffect(() => {
		if (element.current !== null) {
			element.current.addEventListener("click", onClickOutside);
		}

		return () => {
			if (element.current !== null) {
				element.current.removeEventListener("click", onClickOutside);
			}
		};
	}, [onClickOutside]);

	return [setElementRef, clickOutside];

}

export function useWindowWidth() {
	const [width, setWidth] = useState(window.innerWidth);

	const handleResize = useCallback(() => setWidth(window.innerWidth), []);

	useEffect(() => {
		window.addEventListener("resize", handleResize);

		return () => {
			window.removeEventListener("resize", handleResize);
		};
	});

	return width;
}