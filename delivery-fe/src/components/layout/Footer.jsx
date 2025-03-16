import {
  Box,
  Container,
  SimpleGrid,
  Stack,
  Text,
  Heading,
  Input,
  InputGroup,
  InputRightElement,
  Button,
  Divider,
  HStack,
  Icon,
  Link,
  Image,
} from '@chakra-ui/react';
import {
  FaFacebook,
  FaTwitter,
  FaInstagram,
  FaYoutube,
  FaPaperPlane,
  FaMapMarkerAlt,
  FaPhone,
  FaEnvelope,
} from 'react-icons/fa';
import Logo from '../ui/Logo';

const Footer = () => {
  return (
    <Box bg='gray.50' color='gray.700' mt={10}>
      <Container as={Stack} maxW={'container.xl'} py={10}>
        <SimpleGrid
          templateColumns={{ sm: '1fr 1fr', md: '2fr 1fr 1fr 1fr 1fr' }}
          spacing={8}
        >
          <Stack spacing={6}>
            <Box>
              <Logo size='lg' />
            </Box>
            <Text fontSize={'sm'}>
              © 2025 ShopEasy. Tất cả các quyền đã được bảo lưu.
            </Text>
            <Stack spacing={3}>
              <HStack>
                <Icon as={FaMapMarkerAlt} color='gray.600' />
                <Text fontSize='sm'>
                  123 Nguyễn Huệ, Quận 1, TP. Hồ Chí Minh
                </Text>
              </HStack>
              <HStack>
                <Icon as={FaPhone} color='gray.600' />
                <Text fontSize='sm'>1900 1234</Text>
              </HStack>
              <HStack>
                <Icon as={FaEnvelope} color='gray.600' />
                <Text fontSize='sm'>contact@shopeasy.vn</Text>
              </HStack>
            </Stack>
            <HStack spacing={6}>
              <Link href='#' isExternal>
                <Icon as={FaFacebook} w={6} h={6} color='blue.500' />
              </Link>
              <Link href='#' isExternal>
                <Icon as={FaTwitter} w={6} h={6} color='blue.400' />
              </Link>
              <Link href='#' isExternal>
                <Icon as={FaInstagram} w={6} h={6} color='pink.500' />
              </Link>
              <Link href='#' isExternal>
                <Icon as={FaYoutube} w={6} h={6} color='red.500' />
              </Link>
            </HStack>
          </Stack>

          <Stack align={'flex-start'}>
            <Heading as='h5' size='sm' mb={2}>
              Về chúng tôi
            </Heading>
            <Link href={'#'}>Giới thiệu</Link>
            <Link href={'#'}>Tuyển dụng</Link>
            <Link href={'#'}>Tin tức</Link>
            <Link href={'#'}>Liên hệ</Link>
          </Stack>

          <Stack align={'flex-start'}>
            <Heading as='h5' size='sm' mb={2}>
              Hỗ trợ khách hàng
            </Heading>
            <Link href={'#'}>Trung tâm trợ giúp</Link>
            <Link href={'#'}>Chính sách bảo hành</Link>
            <Link href={'#'}>Vận chuyển</Link>
            <Link href={'#'}>Thanh toán</Link>
          </Stack>

          <Stack align={'flex-start'}>
            <Heading as='h5' size='sm' mb={2}>
              Chính sách
            </Heading>
            <Link href={'#'}>Chính sách bảo mật</Link>
            <Link href={'#'}>Điều khoản sử dụng</Link>
            <Link href={'#'}>Chính sách đổi trả</Link>
            <Link href={'#'}>Chính sách vận chuyển</Link>
          </Stack>

          <Stack align={'flex-start'}>
            <Heading as='h5' size='sm' mb={2}>
              Đăng ký nhận tin
            </Heading>
            <Text fontSize={'sm'}>
              Nhận thông tin về sản phẩm mới và khuyến mãi hấp dẫn
            </Text>
            <InputGroup size='md'>
              <Input type={'email'} placeholder={'Email của bạn'} bg='white' />
              <InputRightElement width={'4.5rem'}>
                <Button h={'1.75rem'} size={'sm'} colorScheme='brand' mr={1}>
                  <Icon as={FaPaperPlane} />
                </Button>
              </InputRightElement>
            </InputGroup>
          </Stack>
        </SimpleGrid>

        <Divider my={6} borderColor='gray.300' />

        <Box pt={6}>
          <SimpleGrid columns={{ base: 2, md: 4 }} spacing={4}>
            <Image
              src='https://via.placeholder.com/80x30?text=VISA'
              alt='Visa'
              height='30px'
              objectFit='contain'
            />
            <Image
              src='https://via.placeholder.com/80x30?text=MASTERCARD'
              alt='Mastercard'
              height='30px'
              objectFit='contain'
            />
            <Image
              src='https://via.placeholder.com/80x30?text=PAYPAL'
              alt='PayPal'
              height='30px'
              objectFit='contain'
            />
            <Image
              src='https://via.placeholder.com/80x30?text=MOMO'
              alt='MoMo'
              height='30px'
              objectFit='contain'
            />
          </SimpleGrid>
        </Box>

        <Box pt={6} textAlign='center'>
          <Text fontSize='sm'>
            ShopEasy - Nền tảng mua sắm trực tuyến hàng đầu Việt Nam
          </Text>
          <Text fontSize='xs' mt={2} color='gray.500'>
            Giấy chứng nhận Đăng ký Kinh doanh số 0123456789 do Sở Kế hoạch và
            Đầu tư TP. Hồ Chí Minh cấp ngày 01/01/2025
          </Text>
        </Box>

        <Box pt={4} textAlign='center'>
          <SimpleGrid
            columns={{ base: 2, md: 3 }}
            spacing={4}
            maxW='600px'
            mx='auto'
          >
            <Image
              src='https://via.placeholder.com/120x40?text=BO+CONG+THUONG'
              alt='Bộ Công Thương'
              height='40px'
              objectFit='contain'
            />
            <Image
              src='https://via.placeholder.com/120x40?text=DMCA'
              alt='DMCA Protected'
              height='40px'
              objectFit='contain'
            />
            <Image
              src='https://via.placeholder.com/120x40?text=SSL'
              alt='SSL Secured'
              height='40px'
              objectFit='contain'
            />
          </SimpleGrid>
        </Box>
      </Container>
    </Box>
  );
};

export default Footer;
