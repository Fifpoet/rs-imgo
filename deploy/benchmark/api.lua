request = function()
   -- 随机生成 xyz 的取值
   local z = math.random(1, 4)
   local max = math.pow(2, z) - 1
   local x = math.random(0, max)
   local y = math.random(0, max)

   local path = "/v1/map/"  .. z .. "/" .. x .. "/" .. y

   return wrk.format(nil, path)
end