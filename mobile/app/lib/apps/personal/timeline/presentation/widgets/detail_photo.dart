import 'package:celebut/apps/personal/timeline/presentation/widgets/bottom_widget_components.dart';
import 'package:celebut/apps/personal/timeline/presentation/widgets/detailed_pic_asset.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class DetailPhoto extends StatefulWidget {
  const DetailPhoto({super.key});

  @override
  State<DetailPhoto> createState() => _DetailPhotoState();
}

class _DetailPhotoState extends State<DetailPhoto> {
  int page = 0;
  PageController pageController = PageController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SizedBox(
        height: MediaQuery.of(context).size.height,
        child: Stack(
          children: [
            Container(
              margin: const EdgeInsets.symmetric(horizontal: 5),
              height: MediaQuery.of(context).size.height * 0.82,
              width: double.maxFinite,
              decoration: const BoxDecoration(
                borderRadius: BorderRadius.all(
                  Radius.circular(20),
                ),
                color: Colors.black26,
              ),
              child: PageView(
                controller: pageController,
                onPageChanged: (value) {
                  page = value;
                  setState(() {});
                },
                children: getDetailPictures(pageController),
              ),
            ),
            Positioned(
              top: 30,
              left: 0,
              right: 0,
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  IconButton(
                    onPressed: () {
                      context.pop();
                    },
                    icon: const Icon(
                      Icons.arrow_back_ios,
                      size: 15,
                      color: Colors.black,
                    ),
                  ),
                  IconButton(
                    onPressed: () {},
                    icon: const Icon(
                      Icons.more_vert,
                      size: 19,
                    ),
                  ),
                ],
              ),
            ),
            const Positioned(
              right: 0,
              left: 0,
              bottom: 80,
              child: Padding(
                padding: EdgeInsets.symmetric(horizontal: 20),
                child: BottomWidgetComponents(),
              ),
            ),
          ],
        ),
      ),
    );
  }
}

List<DetailedPicAsset> getDetailPictures(PageController controller) {
  return [
    DetailedPicAsset(
      assetPath:
          'https://images.unsplash.com/photo-1500622944204-b135684e99fd?q=80&w=2922&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
      pageController: controller,
    ),
    DetailedPicAsset(
      assetPath:
          'https://plus.unsplash.com/premium_photo-1673643405538-de0f82933fcb?q=80&w=2942&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
      pageController: controller,
    ),
    DetailedPicAsset(
      assetPath:
          'https://images.unsplash.com/reserve/bOvf94dPRxWu0u3QsPjF_tree.jpg?q=80&w=2952&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
      pageController: controller,
    ),
    DetailedPicAsset(
      assetPath:
          'https://plus.unsplash.com/premium_photo-1673643405538-de0f82933fcb?q=80&w=2942&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
      pageController: controller,
    ),
  ];
}
