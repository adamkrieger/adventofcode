defmodule Day3 do

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

    def toCoords(coords) do
        String.split(coords, ",", trim: true)
        |> (fn parts -> 
            %{
                :x => String.trim(Enum.at(parts, 0)), 
                :y => String.trim(Enum.at(parts, 1))
            } end).()
    end

    def toArea(area) do
        String.split(area, "x", trim: true)
        |> (fn parts ->
            %{
                :xlen => String.trim(Enum.at(parts, 0)),
                :ylen => String.trim(Enum.at(parts, 1))
            } end).()
    end

    def markFabric(next, fabric) do
        #IO.puts(next, fabric)
        # fabric[{next(:xOff),next(:yOff)}] = 1
        new = fabric
        |> Map.update({next.coords.x, next.coords.y}, "once", fn _ -> "twice" end)

        # new
        # |> Map.keys
        # |> Enum.join(",")
        # |> IO.puts()

        new
    end

    def outputSomething(deets) do
        IO.puts("wut")
        IO.puts(deets)
    end
end

result = File.stream!("input/sample.txt")
|> Stream.map(&Day3.splitID/1)
|> Stream.map(&Day3.splitRest/1)
|> Enum.reduce(%{}, &Day3.markFabric/2)
|> Enum.filter(fn {_, v} -> v != "twice" end)
|> Map.new
|> Map.keys()
|> Enum.map(fn {x,y} -> IO.puts(x <> y) end)
|> Stream.run

IO.puts(result)