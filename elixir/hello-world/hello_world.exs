defmodule HelloWorld do

  @spec hello(String.t) :: String.t

  def hello(name \\ nil) do
    case name do
      nil ->
        "Hello, World!"
      _ ->
        "Hello, #{name}!"
    end
  end
end
