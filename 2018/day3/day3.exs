defmodule Day3Parse do

    def splitID(line) do 
        String.split(line, "@", trim: true)
        |> (fn parts -> 
            %{
                :id => String.trim(Enum.at(parts, 0)), 
                :rest => String.trim(Enum.at(parts, 1))
            } end).()
    end

    def splitRest(idRest) do
        String.split(idRest.rest, ":", trim: true)
        |> (fn parts -> 
            %{
                :id => idRest.id, 
                :coords => toCoords(Enum.at(parts, 0)), 
                :area => toArea(Enum.at(parts, 1))
            } end).()
    end

    def trimParseIntOrRaise(strIn) do
        strIn 
        |> String.trim
        |> Integer.parse
        |> (fn 
                {v, _} -> v
                _ -> raise "bad parse - check input"
            end).()
    end

    def toCoords(coords) do
        String.split(coords, ",", trim: true)
        |> (fn parts -> 
            %{
                :x => Enum.at(parts, 0) |> trimParseIntOrRaise, 
                :y => Enum.at(parts, 1) |> trimParseIntOrRaise
                         
            } end).()
    end

    def toArea(area) do
        String.split(area, "x", trim: true)
        |> (fn parts ->
            %{
                :xlen => Enum.at(parts, 0) |> trimParseIntOrRaise,
                :ylen => Enum.at(parts, 1) |> trimParseIntOrRaise
            } end).()
    end

end

defmodule Day3 do

    def expandClaim(claim) do
        #printIDOfClaim(claim,"expandClaim")

        claim
        |> expandXWithColumns
        |> Enum.map(fn {rowStart,restOfClaim} -> {rowStart, restOfClaim.coords.y, restOfClaim.area.ylen} end)
        |> Enum.flat_map(&expandColumnsWithRows/1)
        |> List.flatten
    end

    def expandXWithColumns(claim) do
        #printIDOfClaim(claim, "expandRange")

        Range.new(claim.coords.x, (claim.coords.x + claim.area.xlen - 1))
        |> Enum.map(fn row -> {row, claim} end)
    end

    def expandColumnsWithRows({x,y,ylen}) do 
        [tackOnX(Range.new(y, (y + ylen - 1)), x)]
    end

    def tackOnX(rangeIn, xVal) do
        #IO.puts rangeIn

        rangeIn
        |> Enum.map(fn eachY -> {xVal,eachY} end)
    end

    def markEachSquare(next, fabric) do
        Map.update(fabric, next, "once", fn _ -> "twice" end)
    end

    def markFabric(next, fabric) do
        new = next
        |> Enum.reduce(fabric, &markEachSquare/2)

        new
    end

    def printIDOfClaim(claim, tag) do
        IO.puts tag <> ": " <> claim.id <> " - " <> to_string(claim.coords.x) <> " - " <> to_string(claim.area.ylen)
        claim
    end
end

result = File.stream!("input/day3.input")
|> Stream.map(&Day3Parse.splitID/1)
|> Stream.map(&Day3Parse.splitRest/1)
#|> Stream.map(fn claim -> Day3.printIDOfClaim(claim, "stream") end)
|> Stream.map(&Day3.expandClaim/1)
|> Enum.reduce(%{}, &Day3.markFabric/2)
|> Enum.filter(fn {_, v} -> v == "twice" end)
|> Enum.count
|> (fn len -> IO.puts(to_string(len)) end).()
