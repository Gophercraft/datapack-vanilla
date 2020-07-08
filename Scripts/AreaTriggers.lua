-- Scarlet Monestary Graveyard - Entering
OnAreaTrigger(45, function(plyr)
  if plyr:GetLevel() < 20 then
    plyr:SendRequiredLevelZoneError(20)
    return
  end

  plyr:Teleport(189, 1687.270020, 1050.089966, 18.677299, 1.570800)
end)

-- Deadmines - Entering
OnAreaTrigger(78, function(plyr)
  if plyr:GetLevel() < 10 then
    plyr:SendRequiredLevelZoneError(10)
    return
  end

  plyr:Teleport(36, -14.573200, -385.475006, 62.456100, 1.570800)
end)

-- Stormwind Stockades - Entering
OnAreaTrigger(101, function(plyr)
  if plyr:GetLevel() < 15 then
    plyr:SendRequiredLevelZoneError(15)
    return
  end

  plyr:Teleport(34, 48.984901, 0.483882, -16.394199, 6.283190)
end)

-- Stormwind Vault Entrance
OnAreaTrigger(107, function(plyr)
  plyr:Teleport(35, -0.910000, 40.570000, -24.230000, 0.000000)
end)

-- Stormwind Vault Instance
OnAreaTrigger(109, function(plyr)
  plyr:Teleport(0, -8661.570312, 616.573975, 86.187698, 5.497790)
end)

-- Deadmines - Exiting
OnAreaTrigger(119, function(plyr)
  plyr:Teleport(0, -11208.700195, 1675.900024, 24.573299, 4.712390)
end)

-- Deadmines - Exiting after ship
OnAreaTrigger(121, function(plyr)
  plyr:Teleport(0, -11339.900391, 1572.449951, 94.391602, 1.570800)
end)

-- Shadowfang Keep - Entering
OnAreaTrigger(145, function(plyr)
  if plyr:GetLevel() < 10 then
    plyr:SendRequiredLevelZoneError(10)
    return
  end

  plyr:Teleport(33, -228.190994, 2111.409912, 76.890404, 1.221730)
end)

-- Shadowfang Keep - Exiting
OnAreaTrigger(194, function(plyr)
  plyr:Teleport(0, -233.011002, 1567.500000, 76.892097, 4.276060)
end)

-- The Barrens - Wailing Caverns
OnAreaTrigger(226, function(plyr)
  plyr:Teleport(1, -738.461975, -2217.800049, 16.919001, 6.021390)
end)

-- The Barrens - Wailing Caverns
OnAreaTrigger(228, function(plyr)
  if plyr:GetLevel() < 10 then
    plyr:SendRequiredLevelZoneError(10)
    return
  end

  plyr:Teleport(43, -158.440994, 131.600998, -74.255203, 5.846850)
end)

-- Razorfen Kraul - Exiting
OnAreaTrigger(242, function(plyr)
  plyr:Teleport(1, -4463.319824, -1664.290039, 84.048897, 3.926990)
end)

-- Razorfen Kraul - Entering
OnAreaTrigger(244, function(plyr)
  if plyr:GetLevel() < 15 then
    plyr:SendRequiredLevelZoneError(15)
    return
  end

  plyr:Teleport(47, 1942.270020, 1544.229980, 83.305496, 1.309000)
end)

-- Blackfathom Deeps - Entering
OnAreaTrigger(257, function(plyr)
  if plyr:GetLevel() < 10 then
    plyr:SendRequiredLevelZoneError(10)
    return
  end

  plyr:Teleport(48, -150.233994, 106.594002, -39.778999, 4.450590)
end)

-- Blackfathom Deeps Instance Start
OnAreaTrigger(259, function(plyr)
  plyr:Teleport(1, 4246.680176, 743.401978, -24.857300, 4.712390)
end)

-- Uldaman - Entering
OnAreaTrigger(286, function(plyr)
  if plyr:GetLevel() < 30 then
    plyr:SendRequiredLevelZoneError(30)
    return
  end

  plyr:Teleport(70, -228.858994, 46.101799, -46.018600, 1.570800)
end)

-- Uldaman - Exiting
OnAreaTrigger(288, function(plyr)
  plyr:Teleport(0, -6066.250000, -2954.560059, 209.772003, 3.141590)
end)

-- Gnomeregan Instance Start
OnAreaTrigger(322, function(plyr)
  plyr:Teleport(0, -5162.569824, 927.841003, 257.177002, 4.712390)
end)

-- Gnomeregan - Entering
OnAreaTrigger(324, function(plyr)
  if plyr:GetLevel() < 15 then
    plyr:SendRequiredLevelZoneError(15)
    return
  end

  plyr:Teleport(90, -329.097992, -3.207220, -152.850998, 2.967060)
end)

-- Razorfen Downs Entrance
OnAreaTrigger(442, function(plyr)
  if plyr:GetLevel() < 25 then
    plyr:SendRequiredLevelZoneError(25)
    return
  end

  plyr:Teleport(129, 2593.679932, 1111.229980, 50.951801, 4.712390)
end)

-- Razorfen Downs Instance Start
OnAreaTrigger(444, function(plyr)
  plyr:Teleport(1, -4659.640137, -2524.239990, 81.374001, 0.785398)
end)

-- Altar of Atal'Hakkar (Sunken Temple) - Entering
OnAreaTrigger(446, function(plyr)
  if plyr:GetLevel() < 35 then
    plyr:SendRequiredLevelZoneError(35)
    return
  end

  plyr:Teleport(109, -315.903015, 100.196999, -131.848999, 3.141590)
end)

-- Altar of Atal'Hakkar (Sunken Temple) - Exiting
OnAreaTrigger(448, function(plyr)
  plyr:Teleport(0, -10176.599609, -3995.350098, -112.184998, 3.001970)
end)

-- Stormwind stockades - Exiting
OnAreaTrigger(503, function(plyr)
  plyr:Teleport(0, -8766.110352, 845.499023, 87.995201, 3.839720)
end)

-- Gnomeregan Train Depot Entrance
OnAreaTrigger(523, function(plyr)
  if plyr:GetLevel() < 15 then
    plyr:SendRequiredLevelZoneError(15)
    return
  end

  plyr:Teleport(90, -733.635986, 1.868380, -249.089996, 3.141590)
end)

-- Gnomeregan Train Depot Instance
OnAreaTrigger(525, function(plyr)
  plyr:Teleport(0, -4858.970215, 770.206970, 241.804993, 4.712390)
end)

-- Teddrassil - Ruth Theran
OnAreaTrigger(527, function(plyr)
  plyr:Teleport(1, 8785.790039, 966.982971, 30.199900, 3.403390)
end)

-- Teddrassil - Darnassus
OnAreaTrigger(542, function(plyr)
  plyr:Teleport(1, 9946.250000, 2612.969971, 1316.489990, 4.712390)
end)

-- Scarlet Monestary Graveyard - Exiting
OnAreaTrigger(602, function(plyr)
  plyr:Teleport(0, 2915.340088, -801.580017, 160.332993, 3.490660)
end)

-- Scarlet Monestary Cathedral - Exiting
OnAreaTrigger(604, function(plyr)
  plyr:Teleport(0, 2915.129883, -823.637024, 160.326996, 3.490660)
end)

-- Scarlet Monestary Armory - Exiting
OnAreaTrigger(606, function(plyr)
  plyr:Teleport(0, 2885.959961, -835.802002, 160.326996, 0.349066)
end)

-- Scarlet Monestary Library - Exiting
OnAreaTrigger(608, function(plyr)
  plyr:Teleport(0, 2869.320068, -820.817993, 160.332993, 0.349066)
end)

-- Scarlet Monestary Cathedral - Entering
OnAreaTrigger(610, function(plyr)
  if plyr:GetLevel() < 20 then
    plyr:SendRequiredLevelZoneError(20)
    return
  end

  plyr:Teleport(189, 853.179016, 1319.180054, 18.671400, 1.570800)
end)

-- Scarlet Monestary Armory - Entering
OnAreaTrigger(612, function(plyr)
  if plyr:GetLevel() < 20 then
    plyr:SendRequiredLevelZoneError(20)
    return
  end

  plyr:Teleport(189, 1608.099976, -318.919006, 18.671400, 4.712390)
end)

-- Scarlet Monestary Library - Entering
OnAreaTrigger(614, function(plyr)
  if plyr:GetLevel() < 20 then
    plyr:SendRequiredLevelZoneError(20)
    return
  end

  plyr:Teleport(189, 253.671997, -206.623993, 18.677299, 4.712390)
end)

-- Stormwind - Wizard Sanctum Room
OnAreaTrigger(702, function(plyr)
  plyr:Teleport(0, -9014.940430, 873.325989, 148.615997, 5.497790)
end)

-- Stormwind - Wizard Sanctum Tower Portal
OnAreaTrigger(704, function(plyr)
  plyr:Teleport(0, -9017.459961, 885.901001, 29.620701, 5.497790)
end)

-- Uldaman Instance End
OnAreaTrigger(882, function(plyr)
  plyr:Teleport(0, -6619.970215, -3765.739990, 266.308990, 3.403390)
end)

-- Uldaman Exit
OnAreaTrigger(902, function(plyr)
  if plyr:GetLevel() < 30 then
    plyr:SendRequiredLevelZoneError(30)
    return
  end

  plyr:Teleport(70, -212.949997, 382.427002, -38.748600, 1.396260)
end)

-- Zul'Farrak - Exiting
OnAreaTrigger(922, function(plyr)
  plyr:Teleport(1, -6795.560059, -2890.719971, 8.887420, 3.141590)
end)

-- Zul'Farrak - Entering
OnAreaTrigger(924, function(plyr)
  if plyr:GetLevel() < 35 then
    plyr:SendRequiredLevelZoneError(35)
    return
  end

  plyr:Teleport(209, 1212.670044, 842.039978, 8.933460, 6.283190)
end)

-- Leap of Faith - End of fall
OnAreaTrigger(943, function(plyr)
  plyr:Teleport(1, -5191.270020, -2802.590088, -7.211110, 5.672320)
end)

-- Onyxia's Lair - Exiting
OnAreaTrigger(1064, function(plyr)
  plyr:Teleport(1, -4750.379883, -3754.439941, 49.048500, 0.785398)
end)

-- Blackrock Depths - Entering
OnAreaTrigger(1466, function(plyr)
  if plyr:GetLevel() < 40 then
    plyr:SendRequiredLevelZoneError(40)
    return
  end

  plyr:Teleport(230, 456.928986, 34.092300, -68.089600, 4.712390)
end)

-- Blackrock Spire - Entering
OnAreaTrigger(1468, function(plyr)
  if plyr:GetLevel() < 45 then
    plyr:SendRequiredLevelZoneError(45)
    return
  end

  plyr:Teleport(229, 78.353401, -226.841003, 49.766201, 4.712390)
end)

-- Blackrock Spire - Exiting
OnAreaTrigger(1470, function(plyr)
  plyr:Teleport(0, -7524.700195, -1228.410034, 287.204010, 1.745330)
end)

-- Blackrock Depths - Exiting
OnAreaTrigger(1472, function(plyr)
  plyr:Teleport(0, -7178.410156, -922.151978, 166.091995, 2.007130)
end)

-- Blackrock Spire - Fall out
OnAreaTrigger(2068, function(plyr)
  plyr:Teleport(0, -7558.390137, -1309.430054, 248.453995, 1.570800)
end)

-- Deeprun Tram - Ironforge Instance (Inside)
OnAreaTrigger(2166, function(plyr)
  plyr:Teleport(0, -4839.509766, -1317.739990, 501.868011, 1.483530)
end)

-- Deeprun Tram - Stormwind Instance (Inside)
OnAreaTrigger(2171, function(plyr)
  plyr:Teleport(0, -8354.230469, 524.067993, 91.797096, 2.356190)
end)

-- Deeprun Tram - Stormwind Instance
OnAreaTrigger(2173, function(plyr)
  plyr:Teleport(369, 67.760696, 2490.979980, -4.296490, 3.141590)
end)

-- Deeprun Tram - Ironforge Instance
OnAreaTrigger(2175, function(plyr)
  plyr:Teleport(369, 69.227699, 10.393200, -4.296650, 3.141590)
end)

-- Stratholme - Entering Back Door
OnAreaTrigger(2214, function(plyr)
  if plyr:GetLevel() < 45 then
    plyr:SendRequiredLevelZoneError(45)
    return
  end

  plyr:Teleport(329, 3590.870117, -3643.219971, 138.490997, 5.497790)
end)

-- Stratholme - Entering Left Front
OnAreaTrigger(2216, function(plyr)
  if plyr:GetLevel() < 45 then
    plyr:SendRequiredLevelZoneError(45)
    return
  end

  plyr:Teleport(329, 3392.919922, -3395.030029, 143.134995, 1.570800)
end)

-- Stratholme - Entering Right Front
OnAreaTrigger(2217, function(plyr)
  if plyr:GetLevel() < 45 then
    plyr:SendRequiredLevelZoneError(45)
    return
  end

  plyr:Teleport(329, 3392.840088, -3364.439941, 142.964996, 4.712390)
end)

-- Stratholme - Exiting Back Door
OnAreaTrigger(2221, function(plyr)
  plyr:Teleport(0, 3233.060059, -4048.300049, 108.442001, 2.007130)
end)

-- Ragefire Chasm - Exiting
OnAreaTrigger(2226, function(plyr)
  plyr:Teleport(1, 1814.989990, -4419.229980, -18.815100, 1.919860)
end)

-- Ragefire Chasm - Entering
OnAreaTrigger(2230, function(plyr)
  if plyr:GetLevel() < 8 then
    plyr:SendRequiredLevelZoneError(8)
    return
  end

  plyr:Teleport(389, 0.797643, -8.234290, -15.528800, 4.712390)
end)

-- Hall of Legends - Ogrimmar
OnAreaTrigger(2527, function(plyr)
  plyr:Teleport(450, 222.220001, 74.457001, 25.720900, 0.610865)
end)

-- Hall of Legends - Ogrimmar (Inside)
OnAreaTrigger(2530, function(plyr)
  plyr:Teleport(1, 1637.910034, -4240.250000, 56.174400, 3.926990)
end)

-- Stormwind - Champions Hall
OnAreaTrigger(2532, function(plyr)
  plyr:Teleport(449, -0.401287, 2.400010, -0.255885, 1.570800)
end)

-- Stormwind (Inside) - Champions Hall
OnAreaTrigger(2534, function(plyr)
  plyr:Teleport(0, -8762.820312, 402.433990, 103.901001, 5.497790)
end)

-- Scholomance - Entering
OnAreaTrigger(2567, function(plyr)
  if plyr:GetLevel() < 45 then
    plyr:SendRequiredLevelZoneError(45)
    return
  end

  plyr:Teleport(289, 190.819000, 126.329002, 137.227005, 6.283190)
end)

-- Scholomance - Exiting
OnAreaTrigger(2568, function(plyr)
  plyr:Teleport(0, 1273.910034, -2553.090088, 91.839302, 3.577920)
end)

-- Alterac Valley - Horde Exit
OnAreaTrigger(2606, function(plyr)
  plyr:Teleport(0, 536.494995, -1085.719971, 106.269997, 3.665190)
end)

-- Alterac Valley - Alliance Exit
OnAreaTrigger(2608, function(plyr)
  plyr:Teleport(0, 101.143997, -184.934006, 127.344002, 4.886920)
end)

-- Onyxia's Lair - Entering
OnAreaTrigger(2848, function(plyr)
  if not plyr:HasItem("it:16309") then
    plyr:SendRequiredItemZoneError("it:16309")
    return
  end

  if plyr:GetLevel() < 50 then
    plyr:SendRequiredLevelZoneError(50)
    return
  end

  plyr:Teleport(249, 30.891600, -54.078999, -5.027840, 4.712390)
end)

-- The Molten Bridge
OnAreaTrigger(2886, function(plyr)
  if plyr:GetLevel() < 50 then
    plyr:SendRequiredLevelZoneError(50)
    return
  end

  plyr:Teleport(409, 1091.890015, -466.984985, -105.084000, 3.141590)
end)

-- Molten Core Entrance, Inside
OnAreaTrigger(2890, function(plyr)
  plyr:Teleport(0, -7508.319824, -1039.739990, 180.912003, 3.839720)
end)

-- Maraudon Purple - Exiting
OnAreaTrigger(3126, function(plyr)
  plyr:Teleport(1, -1182.800049, 2877.429932, 85.907997, 1.658060)
end)

-- Maraudon Orange - Exiting
OnAreaTrigger(3131, function(plyr)
  plyr:Teleport(1, -1468.199951, 2614.209961, 76.380402, 6.283190)
end)

-- Maraudon Orange - Entering
OnAreaTrigger(3133, function(plyr)
  if plyr:GetLevel() < 30 then
    plyr:SendRequiredLevelZoneError(30)
    return
  end

  plyr:Teleport(349, 1016.830017, -458.519989, -43.473701, 6.283190)
end)

-- Maraudon Purple - Entering
OnAreaTrigger(3134, function(plyr)
  if plyr:GetLevel() < 30 then
    plyr:SendRequiredLevelZoneError(30)
    return
  end

  plyr:Teleport(349, 755.078003, -617.695984, -32.922199, 1.570800)
end)

-- Dire Maul East - Entering
OnAreaTrigger(3183, function(plyr)
  if plyr:GetLevel() < 45 then
    plyr:SendRequiredLevelZoneError(45)
    return
  end

  plyr:Teleport(429, 47.450100, -153.664993, -2.714390, 5.497790)
end)

-- Dire Maul
OnAreaTrigger(3184, function(plyr)
  if plyr:GetLevel() < 45 then
    plyr:SendRequiredLevelZoneError(45)
    return
  end

  plyr:Teleport(429, -203.166000, -322.997009, -2.724670, 4.712390)
end)

-- Dire Maul East - Entering Back Door
OnAreaTrigger(3185, function(plyr)
  if plyr:GetLevel() < 45 then
    plyr:SendRequiredLevelZoneError(45)
    return
  end

  plyr:Teleport(429, 10.578600, -836.991028, -32.398800, 6.283190)
end)

-- Dire Maul West - Entering
OnAreaTrigger(3186, function(plyr)
  if plyr:GetLevel() < 45 then
    plyr:SendRequiredLevelZoneError(45)
    return
  end

  plyr:Teleport(429, -65.655899, 159.561005, -3.464700, 2.356190)
end)

-- Dire Maul West - Entering Side Door
OnAreaTrigger(3187, function(plyr)
  if plyr:GetLevel() < 45 then
    plyr:SendRequiredLevelZoneError(45)
    return
  end

  plyr:Teleport(429, 33.108299, 158.977005, -3.471260, 0.785398)
end)

-- Dire Maul North - Entering
OnAreaTrigger(3189, function(plyr)
  if plyr:GetLevel() < 45 then
    plyr:SendRequiredLevelZoneError(45)
    return
  end

  plyr:Teleport(429, 254.919998, -19.462999, -2.559600, 5.497790)
end)

-- Dire Maul
OnAreaTrigger(3190, function(plyr)
  plyr:Teleport(1, -3829.679932, 1250.520020, 160.225998, 6.283190)
end)

-- Dire Maul
OnAreaTrigger(3191, function(plyr)
  plyr:Teleport(1, -3749.409912, 1249.250000, 160.220993, 3.141590)
end)

-- Dire Maul
OnAreaTrigger(3193, function(plyr)
  plyr:Teleport(1, -3520.239990, 1078.400024, 161.102997, 1.570800)
end)

-- Dire Maul East - Exiting
OnAreaTrigger(3194, function(plyr)
  plyr:Teleport(1, -3738.620117, 934.521973, 160.975006, 3.141590)
end)

-- Dire Maul
OnAreaTrigger(3195, function(plyr)
  plyr:Teleport(1, -3981.040039, 777.815002, 160.964996, 1.570800)
end)

-- Dire Maul East - Exiting Back Door
OnAreaTrigger(3196, function(plyr)
  plyr:Teleport(1, -4031.250000, 129.345001, 26.474400, 2.705260)
end)

-- Dire Maul East - Exiting after Alzzin
OnAreaTrigger(3197, function(plyr)
  plyr:Teleport(1, -3585.840088, 847.367004, 138.641006, 2.356190)
end)

-- The Molten Core Window Entrance
OnAreaTrigger(3528, function(plyr)
  if plyr:GetLevel() < 50 then
    plyr:SendRequiredLevelZoneError(50)
    return
  end

  if not plyr:QuestDone(7848) then
    plyr:SendRequiredQuestZoneError(7848)
    return
  end

  plyr:Teleport(409, 1091.890015, -466.984985, -105.084000, 3.141590)
end)

-- The Molten Core Window(Lava) Entrance
OnAreaTrigger(3529, function(plyr)
  if plyr:GetLevel() < 50 then
    plyr:SendRequiredLevelZoneError(50)
    return
  end

  if not plyr:QuestDone(7848) then
    plyr:SendRequiredQuestZoneError(7848)
    return
  end

  plyr:Teleport(409, 1091.890015, -466.984985, -105.084000, 3.141590)
end)

-- Blackwing Lair - Entering
OnAreaTrigger(3726, function(plyr)
  if plyr:GetLevel() < 50 then
    plyr:SendRequiredLevelZoneError(50)
    return
  end

  if not plyr:QuestDone(7761) then
    plyr:SendRequiredQuestZoneError(7761)
    return
  end

  plyr:Teleport(469, -7672.319824, -1107.050049, 396.651001, 0.785398)
end)

-- Blackrock Spire, Unknown
OnAreaTrigger(3728, function(plyr)
  plyr:Teleport(0, -7524.700195, -1228.410034, 287.204010, 1.745330)
end)

-- Zul'Gurub - Entering
OnAreaTrigger(3928, function(plyr)
  if plyr:GetLevel() < 50 then
    plyr:SendRequiredLevelZoneError(50)
    return
  end

  plyr:Teleport(309, -11916.599609, -1243.520020, 92.533798, 4.712390)
end)

-- Zul'Gurub Exit
OnAreaTrigger(3930, function(plyr)
  plyr:Teleport(0, -11916.200195, -1206.900024, 92.261200, 1.570800)
end)

-- Arathi Basin Alliance Out
OnAreaTrigger(3948, function(plyr)
  plyr:Teleport(0, -1215.589966, -2531.750000, 21.673401, 3.176500)
end)

-- Arathi Basin Horde Out
OnAreaTrigger(3949, function(plyr)
  plyr:Teleport(0, -831.880981, -3518.520020, 72.483101, 3.368490)
end)

-- Ruins Of Ahn'Qiraj (Inside)
OnAreaTrigger(4006, function(plyr)
  plyr:Teleport(1, -8415.700195, 1502.439941, 30.686199, 5.497790)
end)

-- Ruins of Ahn'Qiraj - Entering
OnAreaTrigger(4008, function(plyr)
  if plyr:GetLevel() < 50 then
    plyr:SendRequiredLevelZoneError(50)
    return
  end

  plyr:Teleport(509, -8436.530273, 1519.170044, 31.907000, 2.617990)
end)

-- Ahn'Qiraj Temple - Entering
OnAreaTrigger(4010, function(plyr)
  if plyr:GetLevel() < 50 then
    plyr:SendRequiredLevelZoneError(50)
    return
  end

  plyr:Teleport(531, -8221.349609, 2014.339966, 129.070999, 0.872665)
end)

-- Ahn'Qiraj Temple - Exiting
OnAreaTrigger(4012, function(plyr)
  plyr:Teleport(1, -8239.009766, 1993.250000, 129.070999, 4.014260)
end)

-- Naxxramas (Exit)
OnAreaTrigger(4055, function(plyr)
  if plyr:GetLevel() < 51 then
    plyr:SendRequiredLevelZoneError(51)
    return
  end

  if not plyr:QuestDone(9378) then
    plyr:SendRequiredQuestZoneError(9378)
    return
  end

  plyr:Teleport(533, 3005.870117, -3435.010010, 293.881989, 0.000000)
end)

-- Naxxramas (Entrance)
OnAreaTrigger(4156, function(plyr)
  if plyr:GetLevel() < 51 then
    plyr:SendRequiredLevelZoneError(51)
    return
  end

  if not plyr:QuestDone(9378) then
    plyr:SendRequiredQuestZoneError(9378)
    return
  end

  plyr:Teleport(533, 3498.280029, -5349.899902, 144.968002, 1.313240)
end)

-- Shadowfang Keep - Falling out
OnAreaTrigger(2406, function(plyr)
  plyr:Teleport(0, -276.240997, 1652.680054, 77.558899, 3.141590)
end)

-- Booty Bay to Gnomeregan
OnAreaTrigger(1103, function(plyr)
  plyr:Teleport(0, -5095.580078, 756.763000, 260.549988, 6.283190)
end)

-- Gnomeregan to Booty Bay
OnAreaTrigger(1104, function(plyr)
  plyr:Teleport(0, -14460.000000, 463.278015, 15.165100, 4.712390)
end)

