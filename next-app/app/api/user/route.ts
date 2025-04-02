import { NextRequest, NextResponse } from "next/server";

export async function GET(request: NextRequest): Promise<NextResponse> {
//   const { searchParams } = new URL(request.url);
//   console.log(searchParams);

  const userApiUrl = `http://go-app:8080/api/user`;

  try {
    const response = await fetch(userApiUrl);
    if (!response.ok) {
      return NextResponse.json(
        { error: "Failed to fetch weather data" },
        { status: response.status }
      );
    }

    const data = await response.json();
    return NextResponse.json(data);
  } catch (error) {
    return NextResponse.json({ error: "Internal Server Error" }, { status: 500 });
  }
}